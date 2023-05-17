package consumer

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"log"
	"post-api/internal/model"
	"sync"
)

type Service interface {
	HandleComment()
	HandlePost()
}

type Handler struct {
	consumer sarama.Consumer
}

func (h *Handler) HandleComment(wg *sync.WaitGroup) {
	partitionConsumer, err := h.consumer.ConsumePartition("comment", 0, sarama.OffsetNewest)

	if err != nil {
		log.Fatal(err)
	}

	for message := range partitionConsumer.Messages() {
		var comment model.CommentDTO
		err := json.Unmarshal(message.Value, &comment)

		if err != nil {
			log.Fatal(err)
		}

		log.Printf("%v", comment)
	}

	wg.Done()
}

func (h *Handler) HandlePost(wg *sync.WaitGroup) {
	partitionConsumer, err := h.consumer.ConsumePartition("post", 0, sarama.OffsetNewest)

	if err != nil {
		log.Fatal(err)
	}

	for message := range partitionConsumer.Messages() {
		var post model.PostDTO
		err := json.Unmarshal(message.Value, &post)

		if err != nil {
			log.Fatal(err)
		}

		log.Printf("%v", post)
	}
	wg.Done()
}

func New(addr string) *Handler {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumer([]string{addr}, config)
	if err != nil {
		log.Fatal(err)
	}

	return &Handler{
		consumer: consumer,
	}
}
