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
	"ws-gateway/internal/handler"
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

	// 1. Luôn lắng nghe NATS Subject riêng của Node Gateway (chat.gateway.{node_id}) để nhận Sender ACK
	natsListener := delivery.NewNATSListener(nc, cfg.Server.NodeID, hub)
	if err := natsListener.Start(ctx); err != nil {
		log.Fatalf("Failed to start NATS delivery listener: %v", err)
	}

	// 2. Nếu chạy mode gRPC, khởi động thêm gRPC Server để nhận Outbound Delivery từ chat-engine
	var grpcListener delivery.DeliveryListener
	if cfg.Server.DeliveryMode == "grpc" {
		grpcListener = delivery.NewGRPCListener(cfg.GRPC.Port, hub)
		if err := grpcListener.Start(ctx); err != nil {
			log.Fatalf("Failed to start gRPC delivery listener: %v", err)
		}
	}

	// HTTP / WebSocket route
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", handler.HandleWebSocket(hub, cfg.Jwt.AccessTokenSecret))
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

	_ = natsListener.Stop(shutdownCtx)
	if grpcListener != nil {
		_ = grpcListener.Stop(shutdownCtx)
	}

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server forced shutdown: %v", err)
	}

	log.Println("WebSocket Gateway exited cleanly")
}
