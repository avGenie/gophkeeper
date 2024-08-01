package main

import (
	"log"

	"github.com/avGenie/gophkeeper/server/internal/app"
	"github.com/avGenie/gophkeeper/server/internal/config"
	"github.com/avGenie/gophkeeper/package/logger"
	"go.uber.org/zap"
)

func main() {
	config, err := config.NewConfig()
	if err != nil {
		log.Fatalf("create config error: %s", err)
	}

	err = logger.Initialize(config.Server.LogLevel, "")
	if err != nil {
		log.Fatalf("initialize logger error: %s", err)
	}

	app, err := app.NewApp(config)
	if err != nil {
		zap.S().Fatal("failed to create app", zap.Error(err))
	}

	app.Run()
}
