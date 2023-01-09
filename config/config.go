package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	HttpPort   int    `env:"HTTP1_PORT" envDefault:"8080"`
	GrpcPort   int    `env:"HTTP2_PORT" envDefault:"9090"`
	DBHost     string `env:"DB_HOST" envDefault:"127.0.0.1"`
	DBPort     int    `env:"DB_PORT" envDefault:"5432"`
	DBUser     string `env:"DB_USER" envDefault:"postgres"`
	DBPassword string `env:"DB_PASSWORD" envDefault:"sa"`
	DBName     string `env:"DB_NAME" envDefault:"template-db"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
