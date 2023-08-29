package main

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "quickstart-events", 0)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	conn.SetWriteDeadline(time.Now().Add(time.Second * 10))
	_, err = conn.WriteMessages(kafka.Message{Value: []byte("hi")})
	if err != nil {
		panic(err)
	}
}
