package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {

	reader := kafka.ReaderConfig{
		Brokers: []string{"kafka:9092"},
		GroupID: "MESSAGE_SERVICE",
		Topic:   "message_request_topic",
	}

	writer := kafka.WriterConfig{
		Brokers:  []string{"kafka:9092"},
		Topic:    "message_response_topic",
		Balancer: &kafka.LeastBytes{},
	}

	conn, err := kafka.Dial("tcp", "localhost:9092")
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	w := kafka.NewWriter(writer)
	r := kafka.NewReader(reader)

	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Println(err)
		}

		switch msg.Key[0] {
		case 0:
			type val struct {
				Id          string    `json:"id"`
				Content     string    `json:"content"`
				Status      int       `json:"status"`
				CreatedAt   time.Time `json:"created_at"`
				ProcessedAt time.Time `json:"processed_at"`
			}

			result := val{}

			if err := json.Unmarshal(msg.Value, &result); err != nil {
				log.Println(err)
			}

			err = w.WriteMessages(context.Background(), kafka.Message{
				Key:     []byte{0},
				Headers: []kafka.Header{{Key: result.Id}},
			})
			if err != nil {
				log.Println(err)
			}
		}
	}
}
