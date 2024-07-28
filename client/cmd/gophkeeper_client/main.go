package main

import (
	"log"

	"github.com/avGenie/gophkeeper/client/internal/app"
	"github.com/avGenie/gophkeeper/client/internal/config"
	"github.com/avGenie/gophkeeper/package/logger"
	"go.uber.org/zap"
)

func main() {
	config := config.NewConfig()

	err := logger.Initialize(config.LogLevel, config.LogPath)
	if err != nil {
		log.Fatalf("initialize logger error: %s", err)
	}

	app, err := app.NewApp(config)
	if err != nil {
		zap.S().Fatal("failed to create app", zap.Error(err))
	}

	app.Run()
}
