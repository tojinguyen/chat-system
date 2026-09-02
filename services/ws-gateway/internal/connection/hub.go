package connection

import (
	"context"
	"log"
	"sync"
	"time"
	"ws-gateway/internal/broker"
	"ws-gateway/internal/config"
	"ws-gateway/internal/domain"
)

type PresenceService interface {
	SetOnline(ctx context.Context, userID, deviceID, gatewayNode string, ttl time.Duration) error
	SetOffline(ctx context.Context, userID, deviceID string) error
	Heartbeat(ctx context.Context, userID, deviceID string, ttl time.Duration) error
}

// Hub maintains the set of active clients and handles broadcasting
type Hub struct {
	// Registered clients: map[userID]map[deviceID]*Client
	clients    map[string]map[string]*Client
	register   chan *Client
	unregister chan *Client
	producer   broker.InboundProducer
	presence   PresenceService
	mu         sync.RWMutex
}

func NewHub(producer broker.InboundProducer, presence PresenceService) *Hub {
	return &Hub{
		clients:    make(map[string]map[string]*Client),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		producer:   producer,
		presence:   presence,
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

			go func(c *Client) {
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()

				if err := h.presence.SetOnline(ctx, c.UserID, c.DeviceID, config.Cfg.Server.NodeID, time.Duration(config.Cfg.Pres.TTL)); err != nil {
					log.Printf("Error setting user online: %v", err)
				}
			}(client)

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

			go func(c *Client) {
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()

				if err := h.presence.SetOffline(ctx, c.UserID, c.DeviceID); err != nil {
					log.Printf("Error setting user offline: %v", err)
				}
			}(client)
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
