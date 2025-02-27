package initializer

import (
	"context"
	"fmt"
	"log"

	kafka "github.com/segmentio/kafka-go"
)

// 	"log"
// 	"plbooking_go_structure1/global"

//	"github.com/segmentio/kafka-go"
//
// )
func InitKafka() {
	log.Println("connecting to kafka")
	brokerAddress := "localhost:29092" // Change this if needed
	topic := "test-topic"

	// Create a new Kafka writer (producer)
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{brokerAddress},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})

	// Test message
	message := kafka.Message{
		Key:   []byte("key"),
		Value: []byte("Hello, Kafka from Segment!"),
	}

	// Send the message
	err := writer.WriteMessages(context.Background(), message)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	fmt.Println("Successfully sent test message to Kafka!")
	writer.Close()
}

// const (
//
//	kafkURL = "localhost:29092"
//	kafkaTopic
//
// )

// func CloseKafka() {
// 	err := global.KafkaProducer.Close()
// 	if err != nil {
// 		log.Fatalf("failed to close kafka producer%v", err)
// 	}
// }
