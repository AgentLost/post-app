package consumer

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"log"
	"storage-service/internal/model"
	"storage-service/internal/service"
)

const (
	postTopic    = "post"
	commentTopic = "comment"
)

type Consumer struct {
	consumer sarama.Consumer
	services *service.Service
}

func (h *Consumer) HandleComment() {
	partitions, err := h.consumer.Partitions(commentTopic)

	if err != nil {
		log.Fatal(err)
	}

	for _, partition := range partitions {
		pc, err := h.consumer.ConsumePartition(commentTopic, partition, sarama.OffsetNewest)
		if err != nil {
			log.Fatal(err)
		}
		go func(pc sarama.PartitionConsumer) {
			for message := range pc.Messages() {
				var post model.PostRequest
				err := json.Unmarshal(message.Value, &post)

				if err != nil {
					log.Fatal(err)
				}

				log.Printf("%v", post)
				err = h.services.PostService.Save(post)

				if err != nil {
					log.Println(err)
				}
			}
		}(pc)
	}
}

func (h *Consumer) HandlePost() {
	partitions, err := h.consumer.Partitions(postTopic)

	if err != nil {
		log.Fatal(err)
	}

	for _, partition := range partitions {
		pc, err := h.consumer.ConsumePartition(postTopic, partition, sarama.OffsetNewest)
		if err != nil {
			log.Fatal(err)
		}
		go func(pc sarama.PartitionConsumer) {
			for message := range pc.Messages() {
				var post model.PostRequest
				err := json.Unmarshal(message.Value, &post)

				if err != nil {
					log.Fatal(err)
				}

				log.Printf("%v", post)
				err = h.services.PostService.Save(post)

				if err != nil {
					log.Println(err)
				}
			}
		}(pc)
	}
}

func New(addr string, services *service.Service) *Consumer {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumer([]string{addr}, config)
	if err != nil {
		log.Fatal(err)
	}

	return &Consumer{
		consumer: consumer,
		services: services,
	}
}
