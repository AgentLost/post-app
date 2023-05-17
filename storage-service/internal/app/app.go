package app

import (
	"log"
	"post-api/config"
	"post-api/internal/consumer"
	"post-api/internal/model"
	"post-api/internal/producer"
	"sync"
	"time"
)

func Run(cfg *config.Config) {

	log.Println("create kafka producer")

	wg := sync.WaitGroup{}
	p := producer.New(cfg.KafkaAddr)
	//comment := model.CommentDTO{Text: "text", PostId: 1, Author: "Me", CreatedAt: time.Now()}
	log.Println("send post.............")
	post := model.PostDTO{Title: "title", Text: "text", Author: "author", CreatedAt: time.Now()}
	wg.Add(1)
	go func() {
		for {
			//err := p.SendComment(comment)
			//if err != nil {
			//	log.Fatal(err)
			//}
			log.Println("send post.............")

			err := p.SendPost(post)

			if err != nil {
				log.Fatal(err)
			}

			log.Println("send.......................")
			time.Sleep(5 * time.Second)
		}
		wg.Done()
	}()

	log.Println("create kafka consumer")
	c := consumer.New(cfg.KafkaAddr)

	wg.Add(1)
	go c.HandlePost(&wg)
	//go h.HandleComment()

	wg.Wait()
}
