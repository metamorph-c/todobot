package config

import (
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

const (
	envLoadError = "Warning: .env file not found, loading from system environment variables"
	envReadError = "failed to load environment variables from .env file: %v"
)

type Config struct {
	// Telegram API token for authenticating requests
	Token string `env:"TELEGRAM_APITOKEN" env-required:"true"`
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println(envLoadError)
	}

	var cnf Config

	if err := cleanenv.ReadEnv(&cnf); err != nil {
		return nil, fmt.Errorf(envReadError, err)
	}

	return &cnf, nil
}
