package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         int
	DBHost       string
	DBPort       int
	DBUser       string
	DBPassword   string
	DBName       string
	JWTSecret    string
	JWTExpiresIn time.Duration
}

func LoadConfig() (*Config, error) {
	// Nạp .env nếu có (không báo lỗi nếu không tìm thấy, vì chạy k8s nạp qua env vars)
	_ = godotenv.Load()

	port := getEnvAsInt("PORT", 3000)
	dbPort := getEnvAsInt("DB_PORT", 5432)

	jwtExpMin := getEnvAsInt("JWT_EXPIRES_MINUTES", 15)
	if envExp := os.Getenv("JWT_EXPIRES_IN"); envExp != "" {
		if d, err := time.ParseDuration(envExp); err == nil {
			jwtExpMin = int(d.Minutes())
		}
	}

	cfg := &Config{
		Port:         port,
		DBHost:       getEnv("DB_HOST", "localhost"),
		DBPort:       dbPort,
		DBUser:       getEnv("DB_USERNAME", "postgres"),
		DBPassword:   getEnv("DB_PASSWORD", "postgrespassword"),
		DBName:       getEnv("DB_DATABASE", "chat_db"),
		JWTSecret:    getEnv("JWT_SECRET", "default_jwt_secret_chat_system"),
		JWTExpiresIn: time.Duration(jwtExpMin) * time.Minute,
	}

	return cfg, nil
}

func (c *Config) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=UTC",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName,
	)
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	valStr := os.Getenv(key)
	if valStr == "" {
		return fallback
	}
	val, err := strconv.Atoi(valStr)
	if err != nil {
		return fallback
	}
	return val
}
