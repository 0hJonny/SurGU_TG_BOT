package bot

import (
	"fmt"
	"strconv"
	"strings"
	"telegram_bot/src/logger"
	"telegram_bot/src/parser"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleCallback(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	log := logger.Logger{}

	log.Info(update.CallbackQuery.Data)

	id, err := strconv.Atoi(strings.TrimPrefix(update.CallbackQuery.Data, "source_"))
	if err != nil {
		log.Warning("Failed to parse source id: " + err.Error())
		return
	}

	rg, err := parser.GetTableListByID(id)
	if err != nil {
		log.Warning("Failed to get table list: " + err.Error())
		return
	}

	b, err := parser.GetTableStructed(rg)
	if err != nil {
		log.Panic("Failed to get table structed: " + err.Error())
	}

	if len(rg) == 0 {
		return
	}

	file := tgbotapi.FileBytes{
		Name:  fmt.Sprintf("%s.txt", rg[0].Group.Name),
		Bytes: b,
	}

	doc := tgbotapi.NewDocument(update.CallbackQuery.Message.Chat.ID, file)
	doc.ReplyToMessageID = update.CallbackQuery.Message.MessageID
	if _, err := bot.Send(doc); err != nil {
		panic(err)
	}
}
