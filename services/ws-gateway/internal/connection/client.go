package connection

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"chat-system/pkg/contracts"
	"ws-gateway/internal/config"
	"ws-gateway/internal/domain"

	"github.com/gorilla/websocket"
)

// Client represents a single active WebSocket connection
type Client struct {
	UserID   string
	DeviceID string
	SendChan chan *domain.WSMessage
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
		var msg domain.WSMessage
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

func (c *Client) handleIncomingMessage(msg *domain.WSMessage) {
	switch msg.Type {
	case domain.WSEventHeartbeat:
		log.Printf("Heartbeat received from user %s", c.UserID)
		go func(c *Client) {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			if err := c.Hub.presence.Heartbeat(ctx, c.UserID, c.DeviceID, time.Duration(config.Cfg.Pres.TTL)); err != nil {
				log.Printf("Error sending heartbeat for user %s: %v", c.UserID, err)
			}
		}(c)
	case domain.WSEventSendMessage:
		brokerMessageType, ok := msg.Type.ToBrokerMessageType()
		if !ok {
			log.Printf("Unhandled message type: %s", msg.Type)
			return
		}

		inboundEvent := contracts.InboundBrokerEvent{
			Type:        brokerMessageType,
			ClientMsgID: msg.ClientMsgID,
			SenderID:    c.UserID,
			DeviceID:    c.DeviceID,
			GatewayNode: config.Cfg.Server.NodeID,
			Payload:     msg.Payload,
			SentAt:      time.Now().UTC(),
		}
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		if err := c.Hub.producer.Publish(ctx, inboundEvent); err != nil {
			c.sendErrorMessage(msg.ClientMsgID, "Failed to send message")
			return
		}
	default:
		log.Printf("Unhandled message type: %s", msg.Type)
	}
}

func (c *Client) sendErrorMessage(clientMsgID string, errorMsg string) {
	errPayload, _ := json.Marshal(domain.FailedToSendPayload{Error: errorMsg})

	errMsg := &domain.WSMessage{
		Type:        domain.WSEventFailedToSend,
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
