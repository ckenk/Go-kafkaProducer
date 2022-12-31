package kenkafka

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func CreateProducer() *kafka.Producer {
	//assuming kafka running locally with default settings
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		panic(err)
	}
	return p
}

func ProduceRecentChange(p *kafka.Producer, recentChange []byte) {

	// Produce messages to a topic (asynchronously)
	topic := "wikimedia.recentchange"
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          recentChange,
	}, nil)

	// Wait for message deliveries
	p.Flush(500)
}

func TerminateProducer(p *kafka.Producer) {
	fmt.Printf("Closing the producer %v", p)
	p.Close()
}
