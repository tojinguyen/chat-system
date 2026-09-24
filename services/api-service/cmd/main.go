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

	"api-service/internal/config"
	deliveryHttp "api-service/internal/delivery/http"
	"api-service/internal/repository"
	"api-service/internal/usecase"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	log.Println("=================================================================")
	log.Println("           🚀 KHỞI ĐỘNG CHAT SYSTEM API SERVICE (GOLANG)        ")
	log.Println("=================================================================")

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("[Config] Failed to load config: %v", err)
	}

	log.Printf("[Config] Connecting to PostgreSQL at %s:%d (DB: %s)...", cfg.DBHost, cfg.DBPort, cfg.DBName)

	gormLogLevel := logger.Warn
	if os.Getenv("GIN_MODE") != "release" {
		gormLogLevel = logger.Info
	}

	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{
		Logger: logger.Default.LogMode(gormLogLevel),
	})
	if err != nil {
		log.Fatalf("[Database] Failed to connect to PostgreSQL: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("[Database] Failed to get underlying sql.DB: %v", err)
	}
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	log.Println("[Database] PostgreSQL connection pool initialized successfully!")

	// 1. Repositories
	userRepo := repository.NewUserRepository(db)
	friendRepo := repository.NewFriendRepository(db)
	convoRepo := repository.NewConversationRepository(db)

	// 2. Usecases
	authUsecase := usecase.NewAuthUsecase(cfg, userRepo)
	userUsecase := usecase.NewUserUsecase(userRepo)
	friendUsecase := usecase.NewFriendUsecase(friendRepo, userRepo)
	convoUsecase := usecase.NewConversationUsecase(convoRepo, userRepo)

	// 3. Handlers
	handlers := &deliveryHttp.Handlers{
		Auth:         deliveryHttp.NewAuthHandler(authUsecase),
		User:         deliveryHttp.NewUserHandler(userUsecase),
		Friend:       deliveryHttp.NewFriendHandler(friendUsecase),
		Conversation: deliveryHttp.NewConversationHandler(convoUsecase),
	}

	// 4. HTTP Router
	router := deliveryHttp.SetupRouter(authUsecase, handlers)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Printf("[HTTP] API Service listening on port %d...", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("[HTTP] Listen error: %v", err)
		}
	}()

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("[HTTP] Shutting down API Service...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("[HTTP] Server forced to shutdown: %v", err)
	}

	_ = sqlDB.Close()
	log.Println("[HTTP] API Service stopped cleanly.")
}
