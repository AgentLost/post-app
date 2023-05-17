package main

import (
	"api-service/config"
	"api-service/internal/app"
	"github.com/caarlos0/env"
	"log"
)

func main() {
	cfg := new(config.Config)

	if err := env.Parse(cfg); err != nil {
		log.Fatal(err)
	}

	app.Run(cfg)
}
