package config

import (
	"fmt"
	"os"
	"time"

	"github.com/caarlos0/env/v11"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Bots        int           `yaml:"bots" env:"SIM_BOTS" envDefault:"50"`
	ConvsPerBot int           `yaml:"convs_per_bot" env:"SIM_CONVS" envDefault:"5"`
	Interval    time.Duration `yaml:"interval" env:"SIM_INTERVAL" envDefault:"200ms"`
	Duration    time.Duration `yaml:"duration" env:"SIM_DURATION" envDefault:"0s"`
	APIURL      string        `yaml:"api_url" env:"SIM_API_URL" envDefault:"http://localhost"`
	WSURL       string        `yaml:"ws_url" env:"SIM_WS_URL" envDefault:"ws://localhost/ws"`
	Password    string        `yaml:"password" env:"SIM_PASSWORD" envDefault:"Pass@123456"`
}

// LoadConfig nạp cấu hình theo thứ tự ưu tiên: Defaults -> File YAML -> Environment Variables
func LoadConfig(path string) (*Config, error) {
	cfg := &Config{}

	// 1. Đọc file YAML nếu file tồn tại
	if path != "" {
		if file, err := os.Open(path); err == nil {
			defer file.Close()
			if err := yaml.NewDecoder(file).Decode(cfg); err != nil {
				return nil, fmt.Errorf("failed to decode yaml config: %w", err)
			}
		}
	}

	// 2. Ghi đè bằng biến môi trường (Environment Variables từ K8s ConfigMap)
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("failed to parse environment variables: %w", err)
	}

	return cfg, nil
}
