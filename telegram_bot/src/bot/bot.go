package bot

import (
	"telegram_bot/src/config"
	"telegram_bot/src/logger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

var log = logger.Logger{}

func NewBot(cfg *config.Config) *Bot {
	bot, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		log.Panic("Failed to create bot: " + err.Error())
	}
	return &Bot{bot: bot}
}

func (b *Bot) Start() {
	log.Info("Starting bot...")

	b.bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := b.bot.GetUpdatesChan(u)
	for update := range updates {
		// Check if we've gotten a message update.
		if update.Message != nil {

			switch update.Message.Command() {
			case "start":
				go handleStart(b.bot, update)
				go handleGetListOld(b.bot, update)
			case "getTableList":
				go handleGetList(b.bot, update)
			}

		} else if update.CallbackQuery != nil {
			// Respond to the callback query, telling Telegram to show the user
			log.Info(update.CallbackQuery.Data)
			go HandleCallback(b.bot, update)
		}
	}
}
