package connection

import (
	"sync"
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
	// TODO: Implement read loop from websocket
}

// WritePump handles pushing messages to the WebSocket connection
func (c *Client) WritePump() {
	// TODO: Implement write loop to websocket
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
