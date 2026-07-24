package config

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server ServerConfig `yaml:"server"`
}

type ServerConfig struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	APIEndpoint string `yaml:"api-endpoint"`
}

func LoadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.Join(errors.New("failed to read config file"), err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, errors.Join(errors.New("failed to unmarshal YAML"), err)
	}

	return &config, nil
}
