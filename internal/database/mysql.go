package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	mysqlDNS = "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
)

type MySQLDatabase struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

func (db *MySQLDatabase) Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf(mysqlDNS, db.User, db.Password, db.Host, db.Port, db.DbName)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
