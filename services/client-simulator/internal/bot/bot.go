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

	conn     *websocket.Conn
	inFlight sync.Map // clientMsgID -> time.Time
	sendChan chan []byte
	stopChan chan struct{}
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
		HandshakeTimeout: 5 * time.Second,
	}

	conn, _, err := dialer.DialContext(ctx, u.String(), nil)
	if err != nil {
		return fmt.Errorf("bot %s dial error: %w", b.Session.Username, err)
	}

	b.conn = conn
	return nil
}

// Start bắt đầu các pump: read, write, heartbeat, và traffic generator
func (b *Bot) Start(ctx context.Context, sendInterval time.Duration) {
	go b.readPump()
	go b.writePump()
	go b.heartbeatPump(ctx)
	go b.trafficPump(ctx, sendInterval)
}

func (b *Bot) Stop() {
	b.closeOnce.Do(func() {
		close(b.stopChan)
		if b.conn != nil {
			_ = b.conn.Close()
		}
	})
}

func (b *Bot) readPump() {
	defer b.Stop()

	type innerPayload struct {
		MessageID   string `json:"message_id"`
		ClientMsgID string `json:"client_msg_id"`
		SenderID    string `json:"sender_id"`
		ReceiverID  string `json:"receiver_id"`
		Type        string `json:"type"`
	}

	for {
		_, data, err := b.conn.ReadMessage()
		if err != nil {
			return
		}

		var env wsEnvelope
		if err := json.Unmarshal(data, &env); err != nil {
			continue
		}

		var inner innerPayload
		if len(env.Payload) > 0 {
			_ = json.Unmarshal(env.Payload, &inner)
		}

		msgID := env.ClientMsgID
		if msgID == "" {
			msgID = inner.ClientMsgID
		}

		// 1. Kiểm tra xem có phải Sender ACK của tin nhắn do chính Bot này gửi đi không
		if msgID != "" {
			if val, ok := b.inFlight.LoadAndDelete(msgID); ok {
				sentAt := val.(time.Time)
				latency := time.Since(sentAt)
				b.Tracker.RecordAck(latency)
				continue
			}
		}

		// 2. Nếu không phải tin nhắn do bot này gửi -> Tin nhắn nhận được từ bạn chat
		if inner.SenderID != "" && inner.SenderID != b.Session.UserID {
			b.Tracker.RecordDelivered()
		} else if env.Type == "SEND_MESSAGE" && msgID == "" {
			b.Tracker.RecordDelivered()
		}
	}
}

func (b *Bot) writePump() {
	ticker := time.NewTicker(20 * time.Second)
	defer func() {
		ticker.Stop()
		b.Stop()
	}()

	for {
		select {
		case <-b.stopChan:
			return
		case msg, ok := <-b.sendChan:
			if !ok {
				return
			}
			if err := b.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				b.Tracker.RecordError()
				return
			}
		case <-ticker.C:
			// Ping định kỳ
			if err := b.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
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

	// Thêm chút jitter ngẫu nhiên để các bot không bắn cùng mili-giây
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
			// Chọn ngẫu nhiên 1 trong các target conversation
			target := b.Conversations[rand.Intn(len(b.Conversations))]

			// Tạo UUIDv7 làm client_msg_id (Time-ordered UUID)
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

			// Ghi nhận thời điểm gửi để đo Latency
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
