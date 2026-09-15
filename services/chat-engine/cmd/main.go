package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	natsclient "chat-system/pkg/nats"
	"chat-system/pkg/telemetry"
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

	log.Printf("Starting Chat Engine worker: %s (workers=%d, buffer=%d)",
		cfg.Worker.ID, cfg.Worker.Workers, cfg.Worker.BufferSize)

	// Initialize Unified Telemetry (Tracing + Profiling + Metrics Server)
	shutdownTelemetry, err := telemetry.Setup(context.Background(), telemetry.SetupConfig{
		ServiceName:      "chat-engine",
		ServiceVersion:   "1.0.0",
		NodeID:           cfg.Worker.ID,
		CollectorTarget:  cfg.Telemetry.CollectorTarget,
		MetricsPort:      cfg.Telemetry.MetricsPort,
		ProfilerServer:   cfg.Profiler.ServerAddress,
		DisableTracing:   cfg.Telemetry.Disabled,
		DisableProfiling: cfg.Profiler.Disabled,
	})
	if err != nil {
		log.Printf("[Warning] Telemetry setup: %v", err)
	}
	defer shutdownTelemetry(context.Background())

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

	// Initialize Inbound Consumer with Partitioned Worker Pool
	inboundConsumer, err := consumer.NewConsumer(consumer.Config{
		NC:          nc,
		Subject:     cfg.NATS.InboundSubject,
		QueueGroup:  cfg.NATS.InboundConsumerGroup,
		NumWorkers:  cfg.Worker.Workers,
		BufferSize:  cfg.Worker.BufferSize,
		ChatUsecase: chatUsecase,
	})
	if err != nil {
		log.Fatalf("Failed to initialize partitioned consumer: %v", err)
	}

	if err := inboundConsumer.Start(ctx); err != nil {
		log.Fatalf("Failed to start inbound consumer: %v", err)
	}

	log.Printf("Consumer started on subject '%s' with queue group '%s' and %d partitioned workers",
		cfg.NATS.InboundSubject, cfg.NATS.InboundConsumerGroup, cfg.Worker.Workers)

	// Wait for termination signal
	<-ctx.Done()

	log.Println("Received termination signal, shutting down Chat Engine...")

	// Graceful shutdown sequence with drain timeout
	drainTimeout := time.Duration(cfg.Worker.DrainTimeoutSeconds) * time.Second
	if drainTimeout <= 0 {
		drainTimeout = 10 * time.Second
	}
	drainCtx, cancel := context.WithTimeout(context.Background(), drainTimeout)
	defer cancel()

	if err := inboundConsumer.Stop(drainCtx); err != nil {
		log.Printf("Error during graceful consumer drain: %v", err)
	}

	log.Println("Chat Engine exited cleanly")
}
