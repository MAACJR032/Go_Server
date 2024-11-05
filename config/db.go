package config

import (
	"API/schemas"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")

	dbPath := "./db/main.db"
	_, err := os.Stat(dbPath)

	// Check if database file exists
	if os.IsNotExist(err) {
		logger.Info("database file not found. creating ...")

		// Create database file and directory
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}

		file.Close()
	}

	// Create DB and connect
	db, err := gorm.Open(sqlite.Open("./db/main.db"), &gorm.Config{})
	if err != nil {
		logger.Errorf("could not open sqlite database: %v", err)
		return nil, err
	}

	// Migrate the Schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("sqlite auto migration error: %v", err)
		return nil, err
	}

	return db, nil
}