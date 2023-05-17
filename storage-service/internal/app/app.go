package app

import (
	"log"
	"storage-service/config"
	"storage-service/internal/consumer"
	"storage-service/internal/db"
	"storage-service/internal/handler"
	"storage-service/internal/service"
)

func Run(cfg *config.Config) {
	Migrate(cfg)
	database := db.New(cfg)

	log.Println("Starting storage-service")
	s := service.New(database)
	h := handler.New(s)
	log.Println("start kafka consumer")
	c := consumer.New(cfg.KafkaAddr, s)

	go c.HandlePost()
	go c.HandleComment()

	log.Fatal(serverRun(cfg.HttpPort, h.InitRoutes()))
}
