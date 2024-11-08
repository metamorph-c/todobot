package config

import (
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

const (
	envLoadError    = "Warning: .env file not found, loading from system environment variables"
	envReadError    = "failed to load environment variables from .env file: %v"
	configReadError = "Warning: db config not found: %v"
)

type Config struct {
	// Telegram API token for authenticating requests
	Token string `env:"TELEGRAM_APITOKEN" env-required:"true"`

	// Path to the database configuration file
	DbPath string `env:"DB_PATH" env-required:"true"`

	Database `yaml:"database"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"db_name"`
	Dialect  string `yaml:"dialect"`
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println(envLoadError)
	}

	var cfg Config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, fmt.Errorf(envReadError, err)
	}

	if err := cleanenv.ReadConfig(cfg.DbPath, &cfg); err != nil {
		log.Printf(configReadError, err)
	}

	return &cfg, nil
}
