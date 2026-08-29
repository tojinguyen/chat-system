package connection

import (
	"log"
	"sync"
	"time"
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
		// TODO: Implement heartbeat handling logic (Reset presence timer, etc.)
	case domain.WSEventSendMessage:
		log.Printf("User %s sent message to conversation", c.UserID)
		// TODO: Implement message routing logic to other users/devices in the conversation
	default:
		log.Printf("Unhandled message type: %s", msg.Type)
	}
}

// WritePump handles pushing messages to the WebSocket connection
func (c *Client) WritePump() {
	for msg := range c.SendChan {
		if err := c.Conn.WriteJSON(msg); err != nil {
			return
		}
	}
}

// Hub maintains the set of active clients and handles broadcasting
type Hub struct {
	// Registered clients: map[userID]map[deviceID]*Client
	clients    map[string]map[string]*Client
	register   chan *Client
	unregister chan *Client
	mu         sync.RWMutex
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[string]map[string]*Client),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			if _, ok := h.clients[client.UserID]; !ok {
				h.clients[client.UserID] = make(map[string]*Client)
			}
			h.clients[client.UserID][client.DeviceID] = client
			h.mu.Unlock()

		case client := <-h.unregister:
			h.mu.Lock()
			if userClients, ok := h.clients[client.UserID]; ok {
				delete(userClients, client.DeviceID)
				if len(userClients) == 0 {
					delete(h.clients, client.UserID)
				}
			}
			close(client.SendChan)
			h.mu.Unlock()
		}
	}
}

// SendToUser pushes a message to all active devices of a user connected to this node
func (h *Hub) SendToUser(userID string, msg *domain.WSMessage) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	if devices, ok := h.clients[userID]; ok {
		for _, client := range devices {
			select {
			case client.SendChan <- msg:
			default:
				// Channel full or blocked
			}
		}
	}
}

func (h *Hub) RegisterClient(client *Client) {
	h.register <- client
}

func (h *Hub) UnregisterClient(client *Client) {
	h.unregister <- client
}
