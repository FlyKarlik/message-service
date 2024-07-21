package main

import (
	"log"

	"github.com/FlyKarlik/message-service/internal/app"
	"github.com/FlyKarlik/message-service/internal/config"
)

func main() {

	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Failed to load configuration: %s", err.Error())
	}

	app := app.New(cfg)

	if err := app.Run(); err != nil {
		log.Fatalf("Failed to run application: %s", err.Error())
	}
}
