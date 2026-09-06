package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	natsclient "chat-system/pkg/nats"
	"chat-worker/internal/config"
	"chat-worker/internal/consumer"
)

func main() {
	cfg, err := config.LoadConfig("configs/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	log.Printf("Starting Chat Engine worker: %s", cfg.Worker.ID)

	// Connect to NATS Broker
	nc, err := natsclient.Connect(cfg.NATS.URL, "chat-worker-"+cfg.Worker.ID)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	defer nc.Close()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Initialize Inbound Consumer with worker queue group
	inboundConsumer := consumer.NewConsumer(
		nc,
		cfg.NATS.InboundSubject,
		cfg.NATS.InboundConsumerGroup,
	)

	if err := inboundConsumer.Start(ctx); err != nil {
		log.Fatalf("Failed to start inbound consumer: %v", err)
	}

	log.Printf("Consumer started on subject '%s' with queue group '%s'",
		cfg.NATS.InboundSubject, cfg.NATS.InboundConsumerGroup)

	// Wait for termination signal
	<-ctx.Done()

	log.Println("Shutting down Chat Engine...")
	if err := inboundConsumer.Stop(); err != nil {
		log.Printf("Error stopping consumer: %v", err)
	}

	log.Println("Chat Engine exited cleanly")
}
