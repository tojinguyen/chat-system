package config

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v11"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Worker    WorkerConfig    `yaml:"worker"`
	NATS      NATSConfig      `yaml:"nats"`
	Database  DatabaseConfig  `yaml:"database"`
	Redis     RedisConfig     `yaml:"redis"`
	Delivery  DeliveryConfig  `yaml:"delivery"`
	Telemetry TelemetryConfig `yaml:"telemetry"`
	Profiler  ProfilerConfig  `yaml:"profiler"`
}

type TelemetryConfig struct {
	CollectorTarget string `yaml:"collector_target" env:"OTEL_COLLECTOR_TARGET" envDefault:"localhost:4317"`
	MetricsPort     int    `yaml:"metrics_port" env:"METRICS_PORT" envDefault:"9092"`
	Disabled        bool   `yaml:"disabled" env:"OTEL_DISABLED" envDefault:"false"`
}

type ProfilerConfig struct {
	ServerAddress string `yaml:"server_address" env:"PYROSCOPE_SERVER" envDefault:"http://localhost:4040"`
	Disabled      bool   `yaml:"disabled" env:"PYROSCOPE_DISABLED" envDefault:"false"`
}

type WorkerConfig struct {
	ID                  string `yaml:"id" env:"WORKER_ID" envDefault:"chat-worker-01"`
	Workers             int    `yaml:"workers" env:"CHAT_WORKER_COUNT" envDefault:"32"`
	BufferSize          int    `yaml:"buffer_size" env:"CHAT_WORKER_BUFFER_SIZE" envDefault:"1024"`
	DrainTimeoutSeconds int    `yaml:"drain_timeout_seconds" env:"CHAT_WORKER_DRAIN_TIMEOUT_SECONDS" envDefault:"10"`
}

type NATSConfig struct {
	URL                   string `yaml:"url" env:"NATS_URL" envDefault:"nats://localhost:4222"`
	InboundSubject        string `yaml:"inbound_subject" env:"NATS_INBOUND_SUBJECT" envDefault:"chat.inbound"`
	InboundStream         string `yaml:"inbound_stream" env:"NATS_INBOUND_STREAM" envDefault:"CHAT_INBOUND"`
	InboundConsumerGroup  string `yaml:"inbound_consumer_group" env:"NATS_INBOUND_CONSUMER_GROUP" envDefault:"chat_workers"`
	OutboundSubjectPrefix string `yaml:"outbound_subject_prefix" env:"NATS_OUTBOUND_SUBJECT_PREFIX" envDefault:"chat.gateway."`
	NotificationSubject   string `yaml:"notification_subject" env:"NATS_NOTIFICATION_SUBJECT" envDefault:"chat.notification"`
}

type DatabaseConfig struct {
	Hosts          []string `yaml:"hosts" env:"CASSANDRA_HOSTS" envSeparator:","`
	Keyspace       string   `yaml:"keyspace" env:"CASSANDRA_KEYSPACE" envDefault:"chat_system"`
	Table          string   `yaml:"table" env:"CASSANDRA_TABLE" envDefault:"messages"`
	TimeoutSeconds int      `yaml:"timeout_seconds" env:"CASSANDRA_TIMEOUT_SECONDS" envDefault:"5"`
}

type RedisConfig struct {
	Addr                  string `yaml:"addr" env:"REDIS_ADDR" envDefault:"localhost:6379"`
	Password              string `yaml:"password" env:"REDIS_PASSWORD"`
	DB                    int    `yaml:"db" env:"REDIS_DB" envDefault:"0"`
	IdempotencyTTLSeconds int    `yaml:"idempotency_ttl_seconds" env:"REDIS_IDEMPOTENCY_TTL" envDefault:"86400"`
}

type DeliveryConfig struct {
	Mode              string `yaml:"mode" env:"DELIVERY_MODE" envDefault:"grpc"` // "grpc" or "broker"
	GRPCPort          int    `yaml:"grpc_port" env:"GRPC_PORT" envDefault:"50051"`
	GRPCServiceSuffix string `yaml:"grpc_service_suffix" env:"GRPC_SERVICE_SUFFIX" envDefault:""`
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
	if c.Worker.ID == "" {
		return fmt.Errorf("worker.id không được để trống")
	}
	if c.NATS.URL == "" {
		return fmt.Errorf("nats.url không được để trống")
	}
	if c.NATS.InboundSubject == "" {
		return fmt.Errorf("nats.inbound_subject không được để trống")
	}
	if len(c.Database.Hosts) == 0 {
		return fmt.Errorf("database.hosts không được để trống")
	}
	if c.Redis.Addr == "" {
		return fmt.Errorf("redis.addr không được để trống")
	}
	return nil
}
