package config

type Config struct {
	NATS NATSConfig `yaml:"nats"`
	FCM  FCMConfig  `yaml:"fcm"`
	APNs APNsConfig `yaml:"apns"`
}

type NATSConfig struct {
	URL                 string `yaml:"url"`
	NotificationSubject string `yaml:"notification_subject"`
	QueueGroup          string `yaml:"queue_group"`
}

type FCMConfig struct {
	CredentialsFile string `yaml:"credentials_file"`
}

type APNsConfig struct {
	KeyFile    string `yaml:"key_file"`
	KeyID      string `yaml:"key_id"`
	TeamID     string `yaml:"team_id"`
	Topic      string `yaml:"topic"`
	Production bool   `yaml:"production"`
}

func LoadConfig(path string) (*Config, error) {
	// TODO: Implement YAML/ENV config loader
	return &Config{
		NATS: NATSConfig{
			URL:                 "nats://localhost:4222",
			NotificationSubject: "chat.notifications",
			QueueGroup:          "notification-worker-group",
		},
	}, nil
}
