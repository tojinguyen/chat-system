package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	natsclient "chat-system/pkg/nats"
	"chat-worker/internal/config"
	"chat-worker/internal/consumer"
	"chat-worker/internal/dispatcher"
	"chat-worker/internal/migrator"
	"chat-worker/internal/presence"
	"chat-worker/internal/repository"
	"chat-worker/internal/usecase"

	"github.com/redis/go-redis/v9"
)

func main() {
	cfg, err := config.LoadConfig("configs/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	log.Printf("Starting Chat Engine worker: %s", cfg.Worker.ID)

	// Run Database Schema Migrations
	if err := migrator.Run(&cfg.Database); err != nil {
		log.Fatalf("Database migration failed: %v", err)
	}

	// Connect to NATS Broker
	nc, err := natsclient.Connect(cfg.NATS.URL, "chat-worker-"+cfg.Worker.ID)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	defer nc.Close()

	// Initialize Cassandra Database Session
	dbSession, err := repository.NewCassandraSession(&cfg.Database)
	if err != nil {
		log.Fatalf("Failed to initialize Cassandra session: %v", err)
	}
	defer dbSession.Close()

	// Initialize Redis Client
	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
	defer redisClient.Close()

	// Initialize Repositories, Presence & Dispatcher
	idempotencyTTL := time.Duration(cfg.Redis.IdempotencyTTLSeconds) * time.Second
	idempotencyRepo := repository.NewRedisIdempotencyRepository(redisClient, idempotencyTTL)
	presenceReader := presence.NewPresenceReader(redisClient)
	messageRepo := repository.NewCassandraMessageRepository(dbSession, cfg.Database.Table)

	eventDispatcher, err := dispatcher.NewEventDispatcher(&cfg.Delivery, nc)
	if err != nil {
		log.Fatalf("Failed to initialize event dispatcher: %v", err)
	}
	defer eventDispatcher.Close()

	// Initialize Usecase
	dbTimeout := time.Duration(cfg.Database.TimeoutSeconds) * time.Second
	chatUsecase := usecase.NewChatUsecase(messageRepo, idempotencyRepo, presenceReader, eventDispatcher, dbTimeout)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Initialize Inbound Consumer with worker queue group
	inboundConsumer := consumer.NewConsumer(
		nc,
		cfg.NATS.InboundSubject,
		cfg.NATS.InboundConsumerGroup,
		chatUsecase,
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
