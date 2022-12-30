# A Go Kafka Producer
It subscribe to the wikimedia change stream and produces Kafka messages.
<br>Work in progress...

This assumes:
- A locally running Kafka instance exists with default settings
- Go runtime has been installed

To run, type in the terminal.

```
go run .
```


If module erros received similar to following when starting:
```
cannot find package "github.com/ckenk/go/kafkaProducer/kenkafka" in any of:
    /usr/local/go/src/github.com/ckenk/go/kafkaProducer/kenkafka (from $GOROOT)
    /home/osboxes/go/src/github.com/ckenk/go/kafkaProducer/kenkafka (from $GOPATH)
cannot find package "github.com/r3labs/sse/v2" in any of:
    /usr/local/go/src/github.com/r3labs/sse/v2 (from $GOROOT)
    /home/osboxes/go/src/github.com/r3labs/sse/v2 (from $GOPATH)
```
enable Go modules in console
```
export GO111MODULE=on
```

Then run again.

A CLI Kafka consumer can be started as:
```
kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic wikimedia.recentchange
```

## To Do
- Add a Kafka Docker image to the project
- Add Go Kafka consumer