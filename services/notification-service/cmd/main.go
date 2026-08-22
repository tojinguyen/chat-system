package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"notification-service/internal/config"
	"notification-service/internal/consumer"
	"notification-service/internal/provider"
	"notification-service/internal/usecase"
)

func main() {
	cfg, err := config.LoadConfig("configs/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	log.Println("Starting Notification Service worker...")

	fcmProvider, _ := provider.NewFCMProvider(cfg.FCM.CredentialsFile)
	apnsProvider, _ := provider.NewAPNsProvider(
		cfg.APNs.KeyFile,
		cfg.APNs.KeyID,
		cfg.APNs.TeamID,
		cfg.APNs.Topic,
		cfg.APNs.Production,
	)

	uc := usecase.NewNotificationUsecase(fcmProvider, apnsProvider)

	notiConsumer, err := consumer.NewNotificationConsumer(
		cfg.NATS.URL,
		cfg.NATS.NotificationSubject,
		cfg.NATS.QueueGroup,
	)
	if err != nil {
		log.Fatalf("Failed to initialize NATS consumer: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		if err := notiConsumer.Start(ctx, uc.HandleNotification); err != nil {
			log.Printf("Notification consumer error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down Notification Service...")
	cancel()
	notiConsumer.Close()
	log.Println("Notification Service exited cleanly")
}
