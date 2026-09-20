package connection

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"chat-system/pkg/contracts"
	"chat-system/pkg/telemetry"
	"ws-gateway/internal/config"
	"ws-gateway/internal/payload"

	"github.com/gorilla/websocket"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// Client represents a single active WebSocket connection
type Client struct {
	UserID   string
	DeviceID string
	SendChan chan *payload.WSMessage
	Conn     *websocket.Conn
	Hub      *Hub
}

// ReadPump handles reading messages from the WebSocket connection
func (c *Client) ReadPump() {
	defer func() {
		c.Hub.UnregisterClient(c)
		c.Conn.Close()
	}()

	c.Conn.SetReadLimit(int64(config.Cfg.Ws.MaxMessageSize))
	pongWait := time.Duration(config.Cfg.Ws.PongWait) * time.Second
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))

	c.Conn.SetPongHandler(func(string) error {
		c.Conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		var msg payload.WSMessage
		err := c.Conn.ReadJSON(&msg)

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("[ReadPump] error reading message from user %s: %v", c.UserID, err)
			}
			break
		}

		c.handleIncomingMessage(&msg)
	}
}

func (c *Client) handleIncomingMessage(msg *payload.WSMessage) {
	switch msg.Type {
	case payload.WSEventHeartbeat:
		// 1. Phản hồi ngay HEARTBEAT_ACK về cho client đo Round-Trip Time (Network Delay)
		ackMsg := &payload.WSMessage{
			Type:        payload.WSEventHeartbeatAck,
			ClientMsgID: msg.ClientMsgID,
			Timestamp:   msg.Timestamp,
		}
		select {
		case c.SendChan <- ackMsg:
		default:
		}

		// 2. Đo Inbound Network Delay và ghi nhận lên Prometheus
		if msg.Timestamp > 0 {
			diff := time.Now().UnixMilli() - msg.Timestamp
			if diff >= 0 && diff < 60000 {
				telemetry.ClientNetworkLatency.WithLabelValues(config.Cfg.Server.NodeID).Observe(float64(diff) / 1000.0)
			}
		}

		// 3. Cập nhật presence bất đồng bộ
		go func(c *Client) {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			if err := c.Hub.presence.Heartbeat(ctx, c.UserID, c.DeviceID, config.Cfg.Server.NodeID, time.Duration(config.Cfg.Pres.TTL)*time.Second); err != nil {
				log.Printf("Error sending heartbeat for user %s: %v", c.UserID, err)
			}
		}(c)
	case payload.WSEventSendMessage:
		brokerMessageType, ok := msg.Type.ToBrokerMessageType()
		if !ok {
			log.Printf("Unhandled message type: %s", msg.Type)
			return
		}

		tracer := telemetry.Tracer("ws-gateway")
		traceCtx, span := tracer.Start(context.Background(), "ws.receive_message",
			trace.WithAttributes(
				attribute.String("client_msg_id", msg.ClientMsgID),
				attribute.String("sender_id", c.UserID),
				attribute.String("device_id", c.DeviceID),
				attribute.String("gateway_node", config.Cfg.Server.NodeID),
			),
		)
		defer span.End()

		inboundEvent := contracts.InboundBrokerEvent{
			Type:        brokerMessageType,
			ClientMsgID: msg.ClientMsgID,
			SenderID:    c.UserID,
			DeviceID:    c.DeviceID,
			GatewayNode: config.Cfg.Server.NodeID,
			Payload:     msg.Payload,
			SentAt:      time.Now().UTC(),
		}
		pubCtx, cancel := context.WithTimeout(traceCtx, 2*time.Second)
		defer cancel()

		if err := c.Hub.producer.Publish(pubCtx, inboundEvent); err != nil {
			c.sendErrorMessage(msg.ClientMsgID, "Failed to send message")
			return
		}
	default:
		log.Printf("Unhandled message type: %s", msg.Type)
	}
}

func (c *Client) sendErrorMessage(clientMsgID string, errorMsg string) {
	errPayload, _ := json.Marshal(payload.FailedToSendPayload{Error: errorMsg})

	errMsg := &payload.WSMessage{
		Type:        payload.WSEventFailedToSend,
		ClientMsgID: clientMsgID,
		Payload:     errPayload,
		Timestamp:   time.Now().UnixMilli(),
	}
	select {
	case c.SendChan <- errMsg:
	default:
		log.Printf("[sendErrorMessage] SendChan full for user %s", c.UserID)
	}
}

// WritePump handles pushing messages to the WebSocket connection
func (c *Client) WritePump() {
	pongWait := time.Duration(config.Cfg.Ws.PongWait) * time.Second
	pingPeriod := (pongWait * 9) / 10
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for {
		select {
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(time.Duration(config.Cfg.Ws.WriteDeadline) * time.Second))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Printf("[WritePump] error sending ping to user %s: %v", c.UserID, err)
				return
			}
		case msg, ok := <-c.SendChan:
			c.Conn.SetWriteDeadline(time.Now().Add(time.Duration(config.Cfg.Ws.WriteDeadline) * time.Second))
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.Conn.WriteJSON(msg); err != nil {
				log.Printf("[WritePump] error writing message to user %s: %v", c.UserID, err)
				return
			}
		}
	}
}
