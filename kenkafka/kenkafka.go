package kenkafka

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func ProduceRecentChange(recentChange []byte) {

	//assuming kafka running locally with default settings
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	// Produce messages to topic (asynchronously)
	topic := "wikimedia.recentchange"
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          recentChange,
	}, nil)

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
}
