package config

import (
	"os"

	"github.com/Somaycon/api-go/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"
	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Info("database file not found, creating...")
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

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errof("sqlite openning error: %v", err)
		return nil, err
	}
	err = db.AutoMigrate(&schemas.Task{}, &schemas.User{})
	if err != nil {
		logger.Errof("sqlite automigration error %v", err)
		return nil, err
	}
	return db, nil
}
