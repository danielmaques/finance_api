package main

import (
	"github.com/danielmaques/finance_api/config"
	"github.com/danielmaques/finance_api/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error %v" ,err)
	}

	router.Initialize()
}