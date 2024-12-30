package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	ServerPort     string `envconfig:"SERVER_PORT" default:"8080"`
	Topic          string `envconfig:"CONSUMER_QUEUE" default:"messages"`
	KakaConnection string `envconfig:"KAFKA_BROKER" default:"kafka:9092"`
}

func LoadConfig() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	return &cfg, err
}
