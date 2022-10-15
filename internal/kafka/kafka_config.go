package kafka

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func StartKafka() {
	conf := kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "someTopic",
		GroupID:  "g1",
		MaxBytes: 10,
	}

	r := kafka.NewReader(conf)

	// Todo: listen for gRPC, and get response data
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("Message is : ", string(m.Value))
	}
}
