package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"chat-system/pkg/contracts"
	natsclient "chat-system/pkg/nats"
	"ws-gateway/internal/config"
	"ws-gateway/internal/connection"
	"ws-gateway/internal/domain"
	"ws-gateway/internal/handler"
	"ws-gateway/internal/presence"
)

func main() {
	cfg, err := config.LoadConfig("configs/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	log.Printf("Starting WebSocket Gateway node: %s on port :%d", cfg.Server.NodeID, cfg.Server.Port)

	// Initialize NATS Connection
	nc, err := natsclient.Connect(cfg.NATS.URL, "ws-gateway-"+cfg.Server.NodeID)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	defer nc.Close()

	// Initialize Generic NATS Inbound Producer
	inboundProducer := natsclient.NewPublisher[contracts.InboundBrokerEvent](nc, cfg.NATS.InboundSubject)

	presenceService := presence.NewPresenceService(cfg.Redis.Addr, cfg.Redis.Password, cfg.Redis.DB)

	hub := connection.NewHub(inboundProducer, presenceService)
	go hub.Run()

	// Initialize Generic NATS Outbound Consumer for this Gateway node
	gatewaySubject := contracts.GatewayNodeSubject(cfg.Server.NodeID)
	outboundConsumer := natsclient.NewSubscriber[contracts.OutboundBrokerEvent](nc, gatewaySubject)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	err = outboundConsumer.Start(ctx, func(ctx context.Context, event contracts.OutboundBrokerEvent) error {
		// Forward message received from Chat Engine down to user connected sockets
		payloadBytes, _ := json.Marshal(event)
		wsMsg := &domain.WSMessage{
			Type:      domain.WSEventSendMessage,
			Timestamp: event.Timestamp,
			Payload:   payloadBytes,
		}
		hub.SendToUser(event.ReceiverID, wsMsg)
		return nil
	})
	if err != nil {
		log.Fatalf("Failed to start outbound consumer on subject %s: %v", gatewaySubject, err)
	}
	log.Printf("Subscribed to outbound topic: %s", gatewaySubject)

	// HTTP / WebSocket route
	http.HandleFunc("/ws", handler.HandleWebSocket(hub, cfg.Jwt.AccessTokenSecret))

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: nil,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Wait for termination signal
	<-ctx.Done()

	log.Println("Shutting down WebSocket Gateway...")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server forced shutdown: %v", err)
	}
	log.Println("WebSocket Gateway exited cleanly")
}
