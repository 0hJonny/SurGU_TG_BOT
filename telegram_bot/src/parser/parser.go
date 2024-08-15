package parser

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"telegram_bot/src/logger"
	"telegram_bot/src/models"
	"time"
)

func GetTableStructed(group models.ResponseGroup) ([]byte, error) {
	var bytesList []byte
	for _, table := range group {

		if len(table.List) == 0 {
			continue
		}

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
	}
	return bytesList, nil
}

func GetTableListByID(id int) (models.ResponseGroup, error) {
	var log = logger.Logger{}

	log.Info("GetTableListByID")

	req, err := http.NewRequest("GET", fmt.Sprintf("https://go.surgu.ru/api/campaigns/40/specialities/%d", id), nil)
	if err != nil {
		log.Panic("Failed to create request: " + err.Error())
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Panic("Failed to get response: " + err.Error())
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panic("Failed to read body: " + err.Error())
	}

	var tableList models.ResponseGroup
	if err := tableList.Parse(body); err != nil {
		log.Panic("Failed to parse body: " + err.Error())
	}

	return tableList, nil
}

func GetTableList() (models.ProfileResponse, []models.ResponseGroup, error) {
	var log = logger.Logger{}

	var tableList models.ProfileResponse
	var groupList []models.ResponseGroup

	log.Info("GetTableList")

	currentTime := time.Now()

	sources, err := GetSources()

	if err != nil {
		log.Panic("Failed to get sources: " + err.Error())
	}

	for _, source := range sources {
		req, err := http.NewRequest("GET", fmt.Sprintf("https://go.surgu.ru/api/campaigns/40/specialities/%d", source.ID), nil)
		if err != nil {
			log.Panic("Failed to create request: " + err.Error())
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Panic("Failed to get response: " + err.Error())
		}

		defer resp.Body.Close()

		dataFolderPath := fmt.Sprintf("data/%s/%s", currentTime.Format("2006-01-02_15:04:05"), fmt.Sprintf("%s-%s", source.Code, source.Name))
		err = os.MkdirAll(dataFolderPath, os.ModePerm)
		if err != nil {
			log.Panic("Failed to create data folder: " + err.Error())
		}

		filePath := fmt.Sprintf("%s/%s.json", dataFolderPath, source.Name)
		file, err := os.Create(filePath)
		if err != nil {
			log.Panic("Failed to create file: " + err.Error())
		}
		defer file.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Panic("Failed to read body: " + err.Error())
		}

		_, err = file.Write(body)
		if err != nil {
			log.Panic("Failed to write file: " + err.Error())
		}

		log.Info(fmt.Sprintf("Saved %s", filePath))

		var rg models.ResponseGroup

		err = rg.Parse(body)

		if err != nil {
			log.Panic("Failed to parse group list: " + err.Error())
		}

		groupList = append(groupList, rg)

		tableList = append(tableList, source)
	}

	return tableList, groupList, nil
}
