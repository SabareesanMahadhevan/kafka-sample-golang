package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func kafkaproducertest() {

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "golangdata",
		Balancer: &kafka.LeastBytes{},
	})

	w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Key-one"),
			Value: []byte("{\r\n\t\"name\": \"Sabareesan Mahadhevan\",\r\n\t\"jobdes\": \"golang_dev\"\r\n}"),
		},
	)

	w.Close()
	fmt.Println("kafka write completed")
}
