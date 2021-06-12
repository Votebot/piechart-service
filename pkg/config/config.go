package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	Dev         bool   `default:"false"`
	BindAddress string `default:":8000" envconfig:"BIND_ADDRESS"`
	SentryDsn string `envconfig:"SENTRY_DSN"`
}

func LoadConfig() *Config {
	_ = godotenv.Load()
	var cfg Config
	err := envconfig.Process("PIECHART", &cfg)
	if err != nil {
		log.Fatal("failed to load config: ", err)
	}
	return &cfg
}