package clients

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	botAPIError = "error creating BotApi instance: %v"
)

func Init(token string) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf(botAPIError, err)
	}

	bot.Debug = true

	return bot, nil
}
