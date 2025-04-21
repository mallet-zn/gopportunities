package main

import (
	"github.com/mallet-zn/gopportunities/config"
	"github.com/mallet-zn/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Config Initialize
	err := config.Init()
	if err != nil {
		logger.Errorf("Config Initialization error: %v", err)
		return
	}

	// Router Initialize
	router.Initialize()
}
