package config

type Config struct {
	Server ServerConfig `yaml:"server"`
	Redis  RedisConfig  `yaml:"redis"`
	NATS   NATSConfig   `yaml:"nats"`
}

type ServerConfig struct {
	Port   int    `yaml:"port"`
	NodeID string `yaml:"node_id"`
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

func LoadConfig(path string) (*Config, error) {
	// TODO: Implement YAML/ENV config loader
	return &Config{
		Server: ServerConfig{
			Port:   8080,
			NodeID: "gateway-node-01",
		},
		NATS: NATSConfig{
			URL:                   "nats://localhost:4222",
			InboundSubject:        "chat.inbound",
			OutboundSubjectPrefix: "chat.gateway.",
		},
	}, nil
}
