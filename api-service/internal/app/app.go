package app

import (
	"api-service/config"
	"api-service/internal/handler"
	"api-service/internal/usecases"
	"log"
)

func Run(cfg *config.Config) {
	uc := usecases.New(cfg.KafkaUrl)
	h := handler.New(cfg.StorageUrl, uc)

	if err := serverRun(cfg.HttpPort, h.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
