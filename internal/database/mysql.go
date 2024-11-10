package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	mysqlDsn       = "%s:%s@tcp(%s:%s)/"
	mysqlDsnWithDb = "%s%s?charset=utf8mb4&parseTime=True&loc=Local"

	// Errors
	connectError = "failed to connect to MySQL server at %s:%s - %v"
)

type MySQLDatabase struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func (mysqlDB *MySQLDatabase) Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf(mysqlDsn, mysqlDB.User, mysqlDB.Password, mysqlDB.Host, mysqlDB.Port)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf(connectError, mysqlDB.Host, mysqlDB.Port, err)
	}

	sql := fmt.Sprintf(createDataBaseIfNotExist, mysqlDB.Name)
	if err := db.Exec(sql).Error; err != nil {
		return nil, fmt.Errorf(failedToCreateDb, err)
	}

	dsnWithDb := fmt.Sprintf(mysqlDsnWithDb, dsn, mysqlDB.Name)
	return gorm.Open(mysql.Open(dsnWithDb), &gorm.Config{})
}

func (mysqlDB *MySQLDatabase) AutoMigrate(db *gorm.DB, models ...interface{}) error {
	return db.AutoMigrate(models...)
}
