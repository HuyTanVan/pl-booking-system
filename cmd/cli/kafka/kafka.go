package test_kafka

// import (
// 	"strings"
// 	"time"

// 	kafka "github.com/segmentio/kafka-go"
// )

// // for producer
// func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
// 	return &kafka.Writer{
// 		Addr:     kafka.TCP(kafkaURL),
// 		Topic:    topic,
// 		Balancer: &kafka.LeastBytes{},
// 	}
// }

// // for comsumer
// func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
// 	brokers := strings.Split(kafkaURL, ",")
// 	return kafka.NewReader(kafka.ReaderConfig{
// 		Brokers:        brokers,
// 		GroupID:        groupID,
// 		Topic:          topic,
// 		MinBytes:       10e3, // 10KB
// 		MaxBytes:       10e6, // 10MB
// 		CommitInterval: time.Second,
// 		StartOffset:    kafka.FirstOffset,
// 	})
// }
