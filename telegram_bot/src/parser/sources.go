package parser

import (
	"fmt"
	"io"
	"net/http"
	"telegram_bot/src/logger"
	"telegram_bot/src/models"
)

func getLink(link string) ([]byte, error) {
	var log = logger.Logger{}
	resp, err := http.Get(link)
	if err != nil {
		log.Panic("Failed to get response: " + err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panic("Failed to read body: " + err.Error())
		return nil, err
	}
	return body, nil
}

func GetSources() (models.ProfileResponse, error) {
	var log = logger.Logger{}
	url := "https://go.surgu.ru/api/campaigns/40/specialities"

	body, err := getLink(url)

	if err != nil {
		log.Panic("Failed to get body: " + err.Error())
		return nil, err
	}

	var profile models.ProfileResponse
	if err := profile.Parse(body); err != nil {
		log.Panic("Failed to parse body: " + err.Error())
		return nil, err
	}

	return profile, nil
}

func GetSource(id int) (models.ProfileResponse, error) {
	var log = logger.Logger{}
	url := fmt.Sprintf("https://go.surgu.ru/api/campaigns/40/specialities/%d", id)

	body, err := getLink(url)

	if err != nil {
		log.Panic("Failed to get body: " + err.Error())
		return models.ProfileResponse{}, err
	}

	var profile models.ProfileResponse
	if err := profile.Parse(body); err != nil {
		log.Panic("Failed to parse body: " + err.Error())
		return models.ProfileResponse{}, err
	}

	return profile, nil
}
