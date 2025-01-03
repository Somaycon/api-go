package main

import (
	"github.com/Somaycon/api-go/config"
	"github.com/Somaycon/api-go/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errof("Config initialization error: %v", err)
		return
	}
	router.Initialize()
}
