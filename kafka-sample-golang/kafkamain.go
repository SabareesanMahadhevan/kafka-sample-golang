package main

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {

	kafkaproducer()
	//kafkaconsumer()
	time.Sleep(10 * time.Minute)
}

func kafkaconsumer() {

	conf := kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "golangdata",
		GroupID:  "test-consumer-group",
		MaxBytes: 18,
	}

	reader := kafka.NewReader(conf)

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("some error occured", err)
			continue
		}
		fmt.Println("Message printing from golang : ", string(m.Key), string(m.Value))
	}

}

//kafka producer to write messages in kafka

func kafkaproducer() {

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "golangdata",
		Balancer: &kafka.LeastBytes{},
	})

	w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: []byte("{\r\n\t\"name\": \"Sabareesan Mahadhevan\",\r\n\t\"test\": \"golang_post\"\r\n}"),
		},
	)

	w.Close()
	fmt.Println("kafka write completed")
}
