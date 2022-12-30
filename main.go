package main

import (
	"github.com/ckenk/go/kafkaProducer/kenkafka"
	"github.com/r3labs/sse/v2"
)

func main() {
	client := sse.NewClient("https://stream.wikimedia.org/v2/stream/recentchange")

	// Subscribe to the wiikimedia `recentchange` SSE stream.
	client.Subscribe("message", func(msg *sse.Event) {
		// Got some data!
		kenkafka.ProduceRecentChange(msg.Data)
	})
}
