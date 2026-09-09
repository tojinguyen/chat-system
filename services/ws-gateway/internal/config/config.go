package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server ServerConfig    `yaml:"server"`
	Redis  RedisConfig     `yaml:"redis"`
	NATS   NATSConfig      `yaml:"nats"`
	Jwt    JwtConfig       `yaml:"jwt"`
	Ws     WebSocketConfig `yaml:"ws"`
	Pres   PresenceConfig  `yaml:"presence"`
	GRPC   GRPCConfig      `yaml:"grpc"`
}

type ServerConfig struct {
	Port         int    `yaml:"port"`
	NodeID       string `yaml:"node_id"`
	DeliveryMode string `yaml:"delivery_mode"` // "grpc" or "broker"
}

type RedisConfig struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type NATSConfig struct {
	URL                   string `yaml:"url"`
	InboundSubject        string `yaml:"inbound_subject"`
	OutboundSubjectPrefix string `yaml:"outbound_subject_prefix"`
}

type JwtConfig struct {
	AccessTokenSecret string `yaml:"access_token_secret"`
	AccessTokenExpiry int64  `yaml:"access_token_expiry"`
}

type WebSocketConfig struct {
	PongWait       int64 `yaml:"pong_wait"`
	MaxMessageSize int   `yaml:"max_message_size"`
	WriteDeadline  int64 `yaml:"write_deadline"`
}

type PresenceConfig struct {
	TTL int64 `yaml:"ttl"`
}

type GRPCConfig struct {
	Port           int    `yaml:"port"`
	AdvertisedAddr string `yaml:"advertised_addr"`
}

var Cfg *Config

func LoadConfig(path string) (*Config, error) {
	// Set default values
	config := &Config{
		Server: ServerConfig{
			Port:         8080,
			NodeID:       "gateway-node-01",
			DeliveryMode: "grpc",
		},
		Redis: RedisConfig{
			Addr: "localhost:6379",
			DB:   0,
		},
		NATS: NATSConfig{
			URL:                   "nats://localhost:4222",
			InboundSubject:        "chat.inbound",
			OutboundSubjectPrefix: "chat.gateway.",
		},
		GRPC: GRPCConfig{
			Port:           50051,
			AdvertisedAddr: "localhost:50051",
		},
		Jwt: JwtConfig{
			AccessTokenSecret: "secret",
			AccessTokenExpiry: 3600,
		},
		Pres: PresenceConfig{
			TTL: 60,
		},
	}

	if path != "" {
		data, err := os.ReadFile(path)
		if err == nil {
			if err := yaml.Unmarshal(data, config); err != nil {
				return nil, fmt.Errorf("failed to parse yaml config: %w", err)
			}
		}
	}

	Cfg = config
	return Cfg, nil
}
