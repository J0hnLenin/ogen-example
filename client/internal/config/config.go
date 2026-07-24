package config

import (
	"errors"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Client ClientConfig `yaml:"client"`
}

type ClientConfig struct {
	ServerURL string        `yaml:"server-url"`
	Interval  time.Duration `yaml:"interval"`
	BatchSize int           `yaml:"batch-size"`
}

func LoadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.Join(errors.New("failed to read config file"), err)
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, errors.Join(errors.New("failed to unmarshal YAML"), err)
	}
	return &cfg, nil
}
