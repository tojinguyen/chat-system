package config

import (
	"fmt"
	"os"

	"go.yaml.in/yaml/v2"
)

type Config struct {
	Worker   WorkerConfig   `yaml:"worker"`
	NATS     NATSConfig     `yaml:"nats"`
	Database DatabaseConfig `yaml:"database"`
	Redis    RedisConfig    `yaml:"redis"`
}

type WorkerConfig struct {
	ID string `yaml:"id"`
}

type NATSConfig struct {
	URL                   string `yaml:"url"`
	InboundSubject        string `yaml:"inbound_subject"`
	InboundStream         string `yaml:"inbound_stream"`
	InboundConsumerGroup  string `yaml:"inbound_consumer_group"`
	OutboundSubjectPrefix string `yaml:"outbound_subject_prefix"`
	NotificationSubject   string `yaml:"notification_subject"`
}

type DatabaseConfig struct {
	Hosts    []string `yaml:"hosts"`
	Keyspace string   `yaml:"keyspace"`
	Table    string   `yaml:"table"`
}

type RedisConfig struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

var Cfg *Config

func LoadConfig(path string) (*Config, error) {
	if path == "" {
		return nil, fmt.Errorf("đường dẫn file config không được để trống")
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("không thể đọc file config tại %s: %w", path, err)
	}

	var config Config

	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("không thể parse file yaml config: %w", err)
	}

	Cfg = &config

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
