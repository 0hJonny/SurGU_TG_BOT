package main

import (
	"telegram_bot/src/bot"
	"telegram_bot/src/config"
	"telegram_bot/src/logger"
)

func main() {
	log := logger.Logger{}

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Panic("Failed to load config: " + err.Error())
	}

	bot.NewBot(cfg).Start()
}
