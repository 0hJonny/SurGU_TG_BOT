package bot

import (
	"fmt"
	"sync"
	"telegram_bot/src/models"
	"telegram_bot/src/parser"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleStart(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	log.Info("handleStart" + update.Message.Text)

	// Пример парсинга и отправки пользователю списка
	// tables := parser.GetTableList()

	msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("Привет, %s.", update.Message.From.UserName))
	// for i, table := range tables {
	// 	msg.Text += "\n" + string(i+1) + ". " + table.Name
	// }

	bot.Send(msg)
}

func handleGetListOld(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	log.Info("handleGetList" + update.Message.Text)

	pr, err := parser.GetSources()

	if err != nil {
		log.Panic("Failed to get sources: " + err.Error())
	}

	keyboards := tgbotapi.InlineKeyboardMarkup{}
	for _, source := range pr {
		keyboards.InlineKeyboard = append(keyboards.InlineKeyboard, tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%s (%s)", source.Name, source.Code), fmt.Sprintf("source_%d", source.ID)),
		))
	}

	// keyboards.InlineKeyboard = append(keyboards.InlineKeyboard, tgbotapi.NewInlineKeyboardRow(
	// 	tgbotapi.NewInlineKeyboardButtonData("Все направления", "source_all"),
	// ))
	msg := tgbotapi.NewMessage(chatID, "Выбери направление:")
	msg.ReplyMarkup = keyboards
	bot.Send(msg)
}

func handleGetList(bot *tgbotapi.BotAPI, update tgbotapi.Update) {

	chatID := update.Message.Chat.ID

	_, gr, err := parser.GetTableList()

	if err != nil {
		log.Panic("Failed to get table list: " + err.Error())
	}

	wg := sync.WaitGroup{}
	wg.Add(len(gr))

	for i, group := range gr {
		go func(i int, group models.ResponseGroup) {
			defer wg.Done()

			for j, table := range group {

				log.Info("handleGetList" + update.Message.Text)

				if len(table.List) == 0 {
					continue
				}

				var bytesList []byte

				// Make table
				list := make([][]string, len(table.List)+1)
				table.List = models.OrderUserProfiles(table.List)
				list[0] = []string{"№", "Имя", "В конкурсе", "Оригинал", "Приоритет", "Сумма", "Тест", "ИД"}
				for k, people := range table.List {
					list[k+1] = []string{
						fmt.Sprintf("%d", k+1),
						people.Identity,
						fmt.Sprintf("%d", people.ToOrder),
						fmt.Sprintf("%d", people.Original),
						fmt.Sprintf("%d", people.Priority),
						fmt.Sprintf("%d", people.ScoresSum),
						fmt.Sprintf("%d", people.ScoresSubjectsSum),
						fmt.Sprintf("%d", people.ScoresAchievementsSum),
					}
				}

				// Format table
				maxLengths := make([]int, len(list[0]))
				for _, row := range list {
					for col, val := range row {
						l := len(val)
						if l > maxLengths[col] {
							maxLengths[col] = l
						}
					}
				}
				for _, row := range list {
					for col, val := range row {
						bytesList = append(bytesList, fmt.Sprintf("%-*s", maxLengths[col], val)...)
						if col != len(row)-1 {
							bytesList = append(bytesList, "| "...)
						}
					}
					bytesList = append(bytesList, "\n"...)
				}

				doc := tgbotapi.NewDocument(chatID, tgbotapi.FileBytes{
					Name:  fmt.Sprintf("%d-%d-(%s).txt", i+1, j+1, table.Group.Name),
					Bytes: bytesList,
				})
				doc.Caption = fmt.Sprintf("%s, (%s)", table.Group.Name, table.Group.EducationCondition.Name)
				bot.Send(doc)
			}
		}(i, group)
	}

	wg.Wait()
}
