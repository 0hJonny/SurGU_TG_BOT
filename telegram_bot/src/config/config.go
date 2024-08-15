package config

import (
	"errors"
	"os"
	"telegram_bot/src/logger"
)

type Config struct {
	Token string
}

func LoadConfig() (*Config, error) {
	cfg := os.Getenv("BOT_TOKEN")
	log := logger.Logger{}

	if cfg == "" {
		log.Panic("BOT_TOKEN is not set")
		return nil, errors.New("BOT_TOKEN is not set")
	}

	log.Info("Loaded BOT_TOKEN!")

	return &Config{
		Token: cfg,
	}, nil
}
