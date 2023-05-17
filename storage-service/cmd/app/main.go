package main

import (
	"github.com/caarlos0/env"
	"log"
	"storage-service/config"
	"storage-service/internal/app"
)

func main() {
	cfg := &config.Config{}
	if err := env.Parse(cfg); err != nil {
		log.Fatal(err)
	}
	log.Printf("%v", cfg)

	app.Run(cfg)
}
