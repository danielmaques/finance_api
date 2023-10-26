package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	db              *gorm.DB
	logger          *Logger
	ConectionString string
	Port            = 0
	SecretKey       []byte
)

func Init() error {
	var err error

	db, err = InitializeSQLite()

	if err != nil {
		return fmt.Errorf("error initializing sqlite: %w", err)
	}

	if err = godotenv.Load(); err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	Port, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return fmt.Errorf("error converting port to int: %w", err)
	}

	ConectionString = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NOME"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))

	return nil
}

func GetSQLiteDB() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {

	logger = NweLogger(p)
	return logger

}
