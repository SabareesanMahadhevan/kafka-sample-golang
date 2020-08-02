package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func kafkaconsumertest() {
	// make a new reader that consumes from topic-A, partition 0, at offset 42
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
		fmt.Println("Message printing from golang : ", string(m.Value))
	}

}
