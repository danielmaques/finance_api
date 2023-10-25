package handler

import (
	"github.com/danielmaques/finance_api/config"
	"gorm.io/gorm"
)

var (
	Logger *config.Logger
	DB     *gorm.DB
)

func InitializeHandler() {
	Logger = config.GetLogger("handler")
	DB = config.GetSQLiteDB()
}
