package bot

import (
	"client-simulator/internal/auth"
	"client-simulator/internal/conversation"
	"client-simulator/internal/metrics"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"sync"
	"sync/atomic"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Bot struct {
	Session       *auth.BotSession
	DeviceID      string
	WSURL         string
	Conversations []conversation.Target
	Tracker       *metrics.Tracker

	connMu      sync.RWMutex
	conn        *websocket.Conn
	isConnected atomic.Bool

	inFlight  sync.Map // clientMsgID -> time.Time
	sendChan  chan []byte
	stopChan  chan struct{}
	closeOnce sync.Once
}

type wsEnvelope struct {
	Type        string          `json:"type"`
	ClientMsgID string          `json:"client_msg_id,omitempty"`
	Timestamp   int64           `json:"timestamp,omitempty"`
	Payload     json.RawMessage `json:"payload,omitempty"`
}

type sendMessagePayload struct {
	ConversationID string `json:"conversation_id"`
	ReceiverID     string `json:"receiver_id"`
	Content        string `json:"content"`
}

func NewBot(session *auth.BotSession, wsURL string, targets []conversation.Target, tracker *metrics.Tracker) *Bot {
	return &Bot{
		Session:       session,
		DeviceID:      fmt.Sprintf("dev_%s", session.Username),
		WSURL:         wsURL,
		Conversations: targets,
		Tracker:       tracker,
		sendChan:      make(chan []byte, 1024),
		stopChan:      make(chan struct{}),
	}
}

// IsConnected kiểm tra bot có đang giữ kết nối WebSocket sống hay không
func (b *Bot) IsConnected() bool {
	return b.isConnected.Load()
}

// Connect thiết lập kết nối WebSocket tới WS Gateway
func (b *Bot) Connect(ctx context.Context) error {
	u, err := url.Parse(b.WSURL)
	if err != nil {
		return fmt.Errorf("invalid ws url: %w", err)
	}

	q := u.Query()
	q.Set("token", b.Session.AccessToken)
	q.Set("device_id", b.DeviceID)
	u.RawQuery = q.Encode()

	dialer := websocket.Dialer{
		HandshakeTimeout: 15 * time.Second,
	}

	conn, _, err := dialer.DialContext(ctx, u.String(), nil)
	if err != nil {
		return fmt.Errorf("bot %s dial error: %w", b.Session.Username, err)
	}

	b.connMu.Lock()
	if b.conn != nil {
		_ = b.conn.Close()
	}
	b.conn = conn
	b.connMu.Unlock()

	b.isConnected.Store(true)
	return nil
}

// Start bắt đầu các pump: read, write, heartbeat và traffic generator
func (b *Bot) Start(ctx context.Context, sendInterval time.Duration) {
	go b.lifecyclePump(ctx)
	go b.writePump(ctx)
	go b.heartbeatPump(ctx)
	go b.trafficPump(ctx, sendInterval)
}

func (b *Bot) Stop() {
	b.closeOnce.Do(func() {
		b.isConnected.Store(false)
		close(b.stopChan)
		b.connMu.Lock()
		if b.conn != nil {
			_ = b.conn.Close()
		}
		b.connMu.Unlock()
	})
}

// lifecyclePump quản lý vòng đời kết nối: đọc tin nhắn và tự động Reconnect khi mất mạng
func (b *Bot) lifecyclePump(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-b.stopChan:
			return
		default:
		}

		if !b.IsConnected() {
			b.reconnect(ctx)
			continue
		}

		b.connMu.RLock()
		conn := b.conn
		b.connMu.RUnlock()

		if conn == nil {
			b.reconnect(ctx)
			continue
		}

		_, data, err := conn.ReadMessage()
		if err != nil {
			// Socket bị đứt (Node Gateway sập hoặc network partition)
			b.isConnected.Store(false)
			b.connMu.Lock()
			if b.conn != nil {
				_ = b.conn.Close()
				b.conn = nil
			}
			b.connMu.Unlock()
			continue
		}

		var env wsEnvelope
		if err := json.Unmarshal(data, &env); err != nil {
			continue
		}

		type innerPayload struct {
			MessageID   string `json:"message_id"`
			ClientMsgID string `json:"client_msg_id"`
			SenderID    string `json:"sender_id"`
			ReceiverID  string `json:"receiver_id"`
			Type        string `json:"type"`
		}

		var inner innerPayload
		if len(env.Payload) > 0 {
			_ = json.Unmarshal(env.Payload, &inner)
		}

		msgID := env.ClientMsgID
		if msgID == "" {
			msgID = inner.ClientMsgID
		}

		// 1. Kiểm tra Sender ACK
		if msgID != "" {
			if val, ok := b.inFlight.LoadAndDelete(msgID); ok {
				sentAt := val.(time.Time)
				latency := time.Since(sentAt)
				b.Tracker.RecordAck(latency)
				continue
			}
		}

		// 2. Tin nhắn nhận được từ bạn chat
		if inner.SenderID != "" && inner.SenderID != b.Session.UserID {
			b.Tracker.RecordDelivered()
		} else if env.Type == "SEND_MESSAGE" && msgID == "" {
			b.Tracker.RecordDelivered()
		}
	}
}

// reconnect thực hiện kết nối lại với Exponential Backoff
func (b *Bot) reconnect(ctx context.Context) {
	backoff := 1 * time.Second
	maxBackoff := 8 * time.Second

	for {
		select {
		case <-ctx.Done():
			return
		case <-b.stopChan:
			return
		case <-time.After(backoff):
		}

		if err := b.Connect(ctx); err == nil {
			// Reconnect thành công!
			b.Tracker.RecordReconnect()
			return
		}

		// Tăng thời gian chờ thử lại (Backoff x 2)
		backoff *= 2
		if backoff > maxBackoff {
			backoff = maxBackoff
		}
	}
}

func (b *Bot) writePump(ctx context.Context) {
	ticker := time.NewTicker(20 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-b.stopChan:
			return
		case msg, ok := <-b.sendChan:
			if !ok {
				return
			}
			if !b.IsConnected() {
				b.Tracker.RecordError()
				continue
			}

			b.connMu.RLock()
			conn := b.conn
			b.connMu.RUnlock()

			if conn != nil {
				if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
					b.isConnected.Store(false)
					b.Tracker.RecordError()
				}
			}
		case <-ticker.C:
			if b.IsConnected() {
				b.connMu.RLock()
				conn := b.conn
				b.connMu.RUnlock()
				if conn != nil {
					_ = conn.WriteMessage(websocket.PingMessage, nil)
				}
			}
		}
	}
}

func (b *Bot) heartbeatPump(ctx context.Context) {
	ticker := time.NewTicker(25 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-b.stopChan:
			return
		case <-ticker.C:
			if !b.IsConnected() {
				continue
			}
			hbEnv := wsEnvelope{
				Type:      "HEARTBEAT",
				Timestamp: time.Now().UnixMilli(),
			}
			bytes, _ := json.Marshal(hbEnv)
			select {
			case b.sendChan <- bytes:
			default:
			}
		}
	}
}

func (b *Bot) trafficPump(ctx context.Context, interval time.Duration) {
	if len(b.Conversations) == 0 || interval <= 0 {
		return
	}

	jitter := time.Duration(rand.Intn(500)) * time.Millisecond
	time.Sleep(jitter)

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-b.stopChan:
			return
		case <-ticker.C:
			if !b.IsConnected() {
				b.Tracker.RecordError()
				continue
			}

			target := b.Conversations[rand.Intn(len(b.Conversations))]

			clientMsgID, err := uuid.NewV7()
			var msgIDStr string
			if err == nil {
				msgIDStr = clientMsgID.String()
			} else {
				msgIDStr = uuid.NewString()
			}

			payloadBytes, _ := json.Marshal(sendMessagePayload{
				ConversationID: target.ConversationID,
				ReceiverID:     target.PartnerID,
				Content:        fmt.Sprintf("Hello from %s to %s at %s", b.Session.Username, target.PartnerID, time.Now().Format("15:04:05.000")),
			})

			env := wsEnvelope{
				Type:        "SEND_MESSAGE",
				ClientMsgID: msgIDStr,
				Timestamp:   time.Now().UnixMilli(),
				Payload:     payloadBytes,
			}

			envBytes, err := json.Marshal(env)
			if err != nil {
				log.Printf("[Bot] Marshal error: %v", err)
				continue
			}

			b.inFlight.Store(msgIDStr, time.Now())
			b.Tracker.RecordSent()

			select {
			case b.sendChan <- envBytes:
			default:
				b.Tracker.RecordError()
			}
		}
	}
}
