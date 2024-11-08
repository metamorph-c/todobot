package database

import (
	"fmt"

	"github.com/metamorph-c/todobot/internal/config"
	"gorm.io/gorm"
)

const (
	mysqlDialect = "mysql"
)

type Database interface {
	Connect() (*gorm.DB, error)
}

func NewDatabase(cfg *config.Database) (Database, error) {
	switch cfg.Dialect {
	case mysqlDialect:
		return &MySQLDatabase{
			Host:     cfg.Host,
			Port:     cfg.Port,
			User:     cfg.User,
			Password: cfg.Password,
			DbName:   cfg.DbName,
		}, nil
	default:
		return nil, fmt.Errorf("unsupported database dialect: %s", cfg.Dialect)
	}
}
