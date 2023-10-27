package config

import (
	"os"

	"github.com/danielmaques/finance_api/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/finance.db"
	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Info("Creating sqlite database")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			logger.Errorf("sqlite initialization error %v", err)
			return nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		defer file.Close()
	}

	db, err := gorm.Open(sqlite.Open("./db/finance.db"), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite initialization error %v", err)
		return nil, err
	}

	// Migrate User
	err = db.AutoMigrate(&schemas.User{})
	if err != nil {
		logger.Errorf("sqlite migration error for users table %v", err)
		return nil, err
	}

	// Migrate Transaction
	err = db.AutoMigrate(&schemas.Transaction{})
	if err != nil {
		logger.Errorf("sqlite migration error for transactions table %v", err)
		return nil, err
	}

	// Migrate House
	err = db.AutoMigrate(&schemas.House{})
	if err != nil {
		logger.Errorf("sqlite migration error for house table %v", err)
		return nil, err
	}

	return db, nil
}
