package main

import (
	"log"

	"github.com/avGenie/gophkeeper/server/internal/app/config"
	"github.com/avGenie/gophkeeper/server/internal/app/logger"
	"github.com/avGenie/gophkeeper/server/internal/app/storage/postgres"
	"go.uber.org/zap"
)

func main() {
	config, err := config.NewConfig()
	if err != nil {
		log.Fatalf("create config error: %s", err)
	}

	err = logger.Initialize(config)
	if err != nil {
		log.Fatalf("initialize logger error: %s", err)
	}

	storage, err := postgres.NewStorage(config.Storage)
	if err != nil {
		zap.L().Fatal("failed to init storage", zap.Error(err))
	}
	defer storage.Close()
}
