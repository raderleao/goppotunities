package main

import (
	"github.com/raderleao/gopportunities/config"
	"github.com/raderleao/gopportunities/router"
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
