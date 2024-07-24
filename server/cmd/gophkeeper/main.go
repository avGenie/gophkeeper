package main

import (
	"fmt"
	"log"

	"github.com/avGenie/gophkeeper/server/internal/app/config"
	"github.com/avGenie/gophkeeper/server/internal/app/logger"
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
}
