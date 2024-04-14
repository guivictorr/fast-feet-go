package main

import (
	"github.com/guivictorr/fast-feet-go/config"
	"github.com/guivictorr/fast-feet-go/router"
)

var logger *config.Logger

func main() {
	logger = config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("error initializing configuration: %v", err)
		return
	}

	router.Initialize()
}
