package main

import (
	"log"

	"github.com/metamorph-c/todobot/internal/config"
	"github.com/metamorph-c/todobot/internal/database"
	"github.com/metamorph-c/todobot/internal/database/models"
)

const (
	configLoadingError = "Error loading config: %v"
	getNewDbError      = "Error getting new database: %v"
	dbConnectionError  = "Database connection error: %v"
	autoMigrateError   = "Auto-migration failed: %v"
	sqlRetriveDbError  = "Failed to retrieve SQL DB from GORM: %v"
)

func getModelsList() []interface{} {
	return []interface{}{
		&models.Task{},
	}
}

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf(configLoadingError, err)
	}

	// Initialize the database
	db, err := database.NewDatabase(&cfg.Database)
	if err != nil {
		log.Fatalf(getNewDbError, err)
	}

	// Connect to the database
	gormDb, err := db.Connect()
	if err != nil {
		log.Fatalf(dbConnectionError, err)
	}

	// Retrieve sql.DB and close the connection pool upon completion
	sqlDB, err := gormDb.DB()
	if err != nil {
		log.Fatalf(sqlRetriveDbError, err)
	}
	defer sqlDB.Close()

	// Perform model migrations
	models := getModelsList()
	if err := db.AutoMigrate(gormDb, models...); err != nil {
		log.Fatalf(autoMigrateError, err)
	}
}
