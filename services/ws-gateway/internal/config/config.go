package config

type Config struct {
	Server ServerConfig    `yaml:"server"`
	Redis  RedisConfig     `yaml:"redis"`
	NATS   NATSConfig      `yaml:"nats"`
	Jwt    JwtConfig       `yaml:"jwt"`
	Ws     WebSocketConfig `yaml:"ws"`
	Pres   PresenceConfig  `yaml:"presence"`
}

type ServerConfig struct {
	Port         int    `yaml:"port"`
	NodeID       string `yaml:"node_id"`
	DeliveryMode string `yaml:"delivery_mode"`
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

var Cfg *Config

func LoadConfig(path string) (*Config, error) {
	Cfg = &Config{
		Server: ServerConfig{
			Port:   8080,
			NodeID: "gateway-node-01",
		},
		NATS: NATSConfig{
			URL:                   "nats://localhost:4222",
			InboundSubject:        "chat.inbound",
			OutboundSubjectPrefix: "chat.gateway.",
		},
		Jwt: JwtConfig{
			AccessTokenSecret: "secret",
			AccessTokenExpiry: 3600,
		},
	}
	return Cfg, nil
}
