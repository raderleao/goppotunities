package main

import (
	"github.com/raderleao/goppotunities.git/config"
	"github.com/raderleao/goppotunities.git/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialeze Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing config: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}
