package config

import (
	"os"

	"github.com/guivictorr/fast-feet-go/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var sqliteLogger *Logger

func createSqliteFile(path string) (*os.File, error) {
	logger.Info("database file not found, creating...")
	err := os.MkdirAll("./db", os.ModePerm)
	if err != nil {
		return nil, err
	}
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func InitializeSQLite() (*gorm.DB, error) {
	sqliteLogger = GetLogger("sqlite")
	dbPath := "./db/main.db"

	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		file, err := createSqliteFile(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.User{}, &schemas.Package{}, &schemas.Recipient{})
	if err != nil {
		logger.Errorf("sqlite migration error: %v", err)
		return nil, err
	}

	return db, nil
}
