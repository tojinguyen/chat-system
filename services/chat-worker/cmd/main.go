package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"chat-worker/internal/config"
	"chat-worker/internal/consumer"
	"chat-worker/internal/dispatcher"
	"chat-worker/internal/presence"
	"chat-worker/internal/repository"
	"chat-worker/internal/usecase"
)

func main() {
	cfg, err := config.LoadConfig("configs/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	log.Printf("Starting Chat Worker daemon: %s", cfg.Worker.ID)

	repo := repository.NewScyllaMessageRepository(cfg.Database.Hosts, cfg.Database.Keyspace)
	pres := presence.NewRedisPresenceChecker(cfg.Redis.Addr, cfg.Redis.Password, cfg.Redis.DB)
	disp := dispatcher.NewNATSDispatcher(cfg.NATS.URL)
	uc := usecase.NewChatWorkerUsecase(repo, pres, disp)

	inboundConsumer, err := consumer.NewInboundConsumer(
		cfg.NATS.URL,
		cfg.NATS.InboundSubject,
		cfg.NATS.InboundConsumerGroup,
	)
	if err != nil {
		log.Fatalf("Failed to initialize NATS consumer: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		if err := inboundConsumer.Start(ctx, uc.ProcessInboundMessage); err != nil {
			log.Printf("Consumer encountered error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down Chat Worker...")
	cancel()
	inboundConsumer.Close()
	log.Println("Chat Worker exited cleanly")
}
