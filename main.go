package main

import "github.com/guivictorr/fast-feet-go/config"

var logger *config.Logger

func main() {
	logger = config.GetLogger("main")
	logger.Info("Hello World")
}
