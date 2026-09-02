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

	"ws-gateway/internal/broker"
	"ws-gateway/internal/config"
	"ws-gateway/internal/connection"
	"ws-gateway/internal/handler"
	"ws-gateway/internal/presence"
)

func main() {
	cfg, err := config.LoadConfig("configs/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	log.Printf("Starting WebSocket Gateway node: %s on port :%d", cfg.Server.NodeID, cfg.Server.Port)

	// Initialize NATS Inbound Producer
	inboundProducer, err := broker.NewInboundProducer(cfg.NATS.URL, cfg.NATS.InboundSubject)
	if err != nil {
		log.Fatalf("Failed to initialize NATS inbound producer: %v", err)
	}
	defer inboundProducer.Close()

	presenceService := presence.NewPresenceService(cfg.Redis.Addr, cfg.Redis.Password, cfg.Redis.DB)

	hub := connection.NewHub(inboundProducer, presenceService)
	go hub.Run()

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

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down WebSocket Gateway...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced shutdown: %v", err)
	}
	log.Println("WebSocket Gateway exited cleanly")
}
