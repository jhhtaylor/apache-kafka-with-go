package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "quickstart-events", 0)
	if err != nil {
		log.Fatalf("Failed to dial leader: %v", err)
	}
	defer conn.Close()

	if err := conn.SetReadDeadline(time.Now().Add(time.Second * 8)); err != nil {
		log.Fatalf("Failed to set read deadline: %v", err)
	}

	batch := conn.ReadBatch(1e3, 1e9)
	if batch == nil {
		log.Fatal("Failed to initiate read batch.")
	}
	defer batch.Close()

	bytes := make([]byte, 1e3)
	for {
		n, err := batch.Read(bytes)
		if err != nil {
			// Check if error is a timeout
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				log.Println("No more messages to consume, exiting gracefully.")
				break
			}
			log.Printf("Error while reading batch: %v", err)
			break
		}
		fmt.Println(string(bytes[:n]))
	}
}
