package database

import (
	"fmt"

	"github.com/metamorph-c/todobot/internal/config"
	"gorm.io/gorm"
)

const (
	// Dialects
	mysqlDialect = "mysql"

	// Errors
	failedToCreateDb       = "failed to create database: %v"
	UsupportedDialectError = "unsupported database dialect: %s"

	// Queries
	createDataBaseIfNotExist = "CREATE DATABASE IF NOT EXISTS %s"
)

type Database interface {
	Connect() (*gorm.DB, error)
	AutoMigrate(*gorm.DB, ...interface{}) error
}

func NewDatabase(cfg *config.Database) (Database, error) {
	switch cfg.Dialect {
	case mysqlDialect:
		return &MySQLDatabase{
			Host:     cfg.Host,
			Port:     cfg.Port,
			User:     cfg.User,
			Password: cfg.Password,
			Name:     cfg.DbName,
		}, nil
	default:
		return nil, fmt.Errorf(UsupportedDialectError, cfg.Dialect)
	}
}
