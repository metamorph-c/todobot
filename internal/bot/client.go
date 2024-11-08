package client

import (
	"fmt"

	"github.com/metamorph-c/todobot/internal/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	botAPIError = "error creating BotApi instance: %v"
)

func Init(cnf *config.Config) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(cnf.Token)
	if err != nil {
		return nil, fmt.Errorf(botAPIError, err)
	}

	bot.Debug = true

	return bot, nil
}
