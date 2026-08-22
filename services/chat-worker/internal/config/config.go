package config

type Config struct {
	Worker   WorkerConfig   `yaml:"worker"`
	NATS     NATSConfig     `yaml:"nats"`
	Database DatabaseConfig `yaml:"database"`
	Redis    RedisConfig    `yaml:"redis"`
}

type WorkerConfig struct {
	ID          string `yaml:"id"`
	Concurrency int    `yaml:"concurrency"`
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
	Addr                  string `yaml:"addr"`
	Password              string `yaml:"password"`
	DB                    int    `yaml:"db"`
	IdempotencyTTLSeconds int    `yaml:"idempotency_ttl_seconds"`
}

func LoadConfig(path string) (*Config, error) {
	// TODO: Implement YAML/ENV config loader
	return &Config{
		Worker: WorkerConfig{
			ID:          "chat-worker-01",
			Concurrency: 10,
		},
		NATS: NATSConfig{
			URL:                   "nats://localhost:4222",
			InboundSubject:        "chat.inbound",
			InboundStream:         "CHAT_INBOUND",
			InboundConsumerGroup:  "chat-worker-group",
			OutboundSubjectPrefix: "chat.gateway.",
			NotificationSubject:   "chat.notifications",
		},
	}, nil
}
