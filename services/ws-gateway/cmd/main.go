package main

import (
	"context"
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
	"ws-gateway/internal/delivery"
	"ws-gateway/internal/presence"
)

func main() {
	cfg, err := config.LoadConfig("configs/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	log.Printf("Starting WebSocket Gateway node: %s (mode: %s, ws_port: :%d)",
		cfg.Server.NodeID, cfg.Server.DeliveryMode, cfg.Server.Port)

	// Initialize NATS Connection for Inbound events
	nc, err := natsclient.Connect(cfg.NATS.URL, "ws-gateway-"+cfg.Server.NodeID)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	defer nc.Close()

	// Initialize Generic NATS Inbound Producer
	inboundProducer := natsclient.NewPublisher[contracts.InboundBrokerEvent](nc, cfg.NATS.InboundSubject)

	// Initialize Presence Service
	presenceService := presence.NewPresenceService(cfg.Redis.Addr, cfg.Redis.Password, cfg.Redis.DB)

	// Initialize Connection Hub
	hub := connection.NewHub(inboundProducer, presenceService)
	go hub.Run()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Initialize Outbound Delivery Listener based on configured Mode
	var listener delivery.DeliveryListener
	switch cfg.Server.DeliveryMode {
	case "grpc":
		listener = delivery.NewGRPCListener(cfg.GRPC.Port, hub)
	case "broker":
		listener = delivery.NewNATSListener(nc, cfg.Server.NodeID, hub)
	default:
		log.Fatalf("Unsupported delivery mode '%s'. Must be 'grpc' or 'broker'", cfg.Server.DeliveryMode)
	}

	if err := listener.Start(ctx); err != nil {
		log.Fatalf("Failed to start %s delivery listener: %v", cfg.Server.DeliveryMode, err)
	}

	// HTTP / WebSocket route
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", delivery.HandleWebSocket(hub, cfg.Jwt.AccessTokenSecret))
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: mux,
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

	if err := listener.Stop(shutdownCtx); err != nil {
		log.Printf("Error stopping delivery listener: %v", err)
	}

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server forced shutdown: %v", err)
	}

	log.Println("WebSocket Gateway exited cleanly")
}
