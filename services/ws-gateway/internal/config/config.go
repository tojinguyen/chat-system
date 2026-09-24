package config

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v11"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Server    ServerConfig    `yaml:"server"`
	Redis     RedisConfig     `yaml:"redis"`
	NATS      NATSConfig      `yaml:"nats"`
	Jwt       JwtConfig       `yaml:"jwt"`
	Ws        WebSocketConfig `yaml:"ws"`
	Pres      PresenceConfig  `yaml:"presence"`
	GRPC      GRPCConfig      `yaml:"grpc"`
	Telemetry TelemetryConfig `yaml:"telemetry"`
	Profiler  ProfilerConfig  `yaml:"profiler"`
}

type TelemetryConfig struct {
	CollectorTarget string `yaml:"collector_target" env:"OTEL_COLLECTOR_TARGET" envDefault:"localhost:4317"`
	MetricsPort     int    `yaml:"metrics_port" env:"METRICS_PORT" envDefault:"9091"`
	Disabled        bool   `yaml:"disabled" env:"OTEL_DISABLED" envDefault:"false"`
}

type ProfilerConfig struct {
	ServerAddress string `yaml:"server_address" env:"PYROSCOPE_SERVER" envDefault:"http://localhost:4040"`
	Disabled      bool   `yaml:"disabled" env:"PYROSCOPE_DISABLED" envDefault:"false"`
}

type ServerConfig struct {
	Port         int    `yaml:"port" env:"WS_PORT" envDefault:"8080"`
	NodeID       string `yaml:"node_id" env:"SERVER_NODE_ID" envDefault:"gateway-node-01"`
	DeliveryMode string `yaml:"delivery_mode" env:"DELIVERY_MODE" envDefault:"grpc"` // "grpc" or "broker"
}

type RedisConfig struct {
	Addr     string `yaml:"addr" env:"REDIS_ADDR" envDefault:"localhost:6379"`
	Password string `yaml:"password" env:"REDIS_PASSWORD"`
	DB       int    `yaml:"db" env:"REDIS_DB" envDefault:"0"`
}

type NATSConfig struct {
	URL                   string `yaml:"url" env:"NATS_URL" envDefault:"nats://localhost:4222"`
	InboundSubject        string `yaml:"inbound_subject" env:"NATS_INBOUND_SUBJECT" envDefault:"chat.inbound"`
	OutboundSubjectPrefix string `yaml:"outbound_subject_prefix" env:"NATS_OUTBOUND_SUBJECT_PREFIX" envDefault:"chat.gateway."`
}

type JwtConfig struct {
	AccessTokenSecret string `yaml:"access_token_secret" env:"JWT_ACCESS_TOKEN_SECRET" envDefault:"secret"`
	AccessTokenExpiry int64  `yaml:"access_token_expiry" env:"JWT_ACCESS_TOKEN_EXPIRY" envDefault:"3600"`
}

type WebSocketConfig struct {
	PongWait       int64 `yaml:"pong_wait" env:"WS_PONG_WAIT" envDefault:"60"`
	MaxMessageSize int   `yaml:"max_message_size" env:"WS_MAX_MESSAGE_SIZE" envDefault:"4096"`
	WriteDeadline  int64 `yaml:"write_deadline" env:"WS_WRITE_DEADLINE" envDefault:"10"`
	SendBufferSize int   `yaml:"send_buffer_size" env:"WS_SEND_BUFFER_SIZE" envDefault:"32"`
}

type PresenceConfig struct {
	TTL int64 `yaml:"ttl" env:"PRESENCE_TTL" envDefault:"60"`
}

type GRPCConfig struct {
	Port           int    `yaml:"port" env:"GRPC_PORT" envDefault:"50051"`
	AdvertisedAddr string `yaml:"advertised_addr" env:"GRPC_ADVERTISED_ADDR" envDefault:"localhost:50051"`
}

var Cfg *Config

func LoadConfig(path string) (*Config, error) {
	var cfg Config

	// 1. Đọc file YAML nếu có (làm baseline/default)
	if path != "" {
		data, err := os.ReadFile(path)
		if err == nil {
			if err := yaml.Unmarshal(data, &cfg); err != nil {
				return nil, fmt.Errorf("không thể parse file yaml config: %w", err)
			}
		}
	}

	// 2. Tự động đọc và ghi đè từ Environment Variables (Type-Safe qua caarlos0/env)
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("không thể parse biến môi trường: %w", err)
	}

	Cfg = &cfg

	if err := Cfg.Validate(); err != nil {
		return nil, fmt.Errorf("cấu hình không hợp lệ: %w", err)
	}

	return Cfg, nil
}

func (c *Config) Validate() error {
	if c.NATS.URL == "" {
		return fmt.Errorf("nats.url không được để trống")
	}
	if c.Redis.Addr == "" {
		return fmt.Errorf("redis.addr không được để trống")
	}
	if c.Jwt.AccessTokenSecret == "" {
		return fmt.Errorf("jwt.access_token_secret không được để trống")
	}
	return nil
}
