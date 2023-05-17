package usecases

import (
	"api-service/internal/model"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
)

type UseCases struct {
	syncProducer sarama.SyncProducer
}

func (uc *UseCases) SendComment(request model.CommentRequest) error {
	jsonMessage, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	msg := &sarama.ProducerMessage{
		Topic: "comment",
		Value: sarama.StringEncoder(jsonMessage),
	}
	partition, offset, err := uc.syncProducer.SendMessage(msg)
	if err != nil {
		log.Fatalf("Error sending message to Kafka: %s", err.Error())
	}
	fmt.Printf("Message sent to partition %d at offset %d\n", partition, offset)
	return nil
}

func (uc *UseCases) SendPost(request model.PostRequest) error {
	jsonMessage, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	msg := &sarama.ProducerMessage{
		Topic: "post",
		Value: sarama.StringEncoder(jsonMessage),
	}
	partition, offset, err := uc.syncProducer.SendMessage(msg)
	if err != nil {
		log.Fatalf("Error sending message to Kafka: %s", err.Error())
	}
	fmt.Printf("Message sent to partition %d at offset %d\n", partition, offset)
	return nil
}

func New(addr string) *UseCases {
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer([]string{addr}, cfg)

	if err != nil {
		log.Fatal(err)
	}

	return &UseCases{
		syncProducer: producer,
	}
}
