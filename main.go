package main

import (
	"API/config"
	"API/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Debugf("config initialization error: %v", err)
		return
	}

	router.Initialize()
}