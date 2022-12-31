package main

import (
	"fmt"

	"github.com/ckenk/go/kafkaProducer/kenkafka"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/r3labs/sse/v2"
)

func main() {
	client := sse.NewClient("https://stream.wikimedia.org/v2/stream/recentchange")

	p := kenkafka.CreateProducer()
	// defer p.Close()
	defer kenkafka.TerminateProducer(p)

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

	// Subscribe to the wiikimedia `recentchange` SSE stream.
	client.Subscribe("message", func(msg *sse.Event) {
		// Got some data!
		kenkafka.ProduceRecentChange(p, msg.Data)
	})

}
