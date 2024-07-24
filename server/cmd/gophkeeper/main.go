package main

import (
	"log"

	"github.com/avGenie/gophkeeper/server/internal/app/config"
)

func main() {
	_, err := config.NewConfig()
	if err != nil {
		log.Fatalf("config error: %s", err)
	}
}
