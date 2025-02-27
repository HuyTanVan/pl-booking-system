package kafka_service

import (
	"context"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

//	type IKafkaProducer interface {
//		Connect(topic string, partition int) (*kafka.Conn, error)
//		WriteMessage(s string)
//	}
type KafkaProducer struct {
	kwriter *kafka.Writer
}

func NewKafkaProducer(config Config, pConfig ProducerConfig) *KafkaProducer {
	return &KafkaProducer{
		kwriter: kafka.NewWriter(kafka.WriterConfig{
			Brokers:      config.Brokers,
			Topic:        config.Topic,
			Balancer:     &kafka.LeastBytes{},
			BatchSize:    pConfig.BatchSize,
			BatchTimeout: time.Duration(pConfig.BatchTimeout) * time.Millisecond,
			RequiredAcks: pConfig.RequiredAcks,
			Async:        pConfig.Async,
		}),
	}

}

// SendMessage sends a message to the Kafka topic.
func (p *KafkaProducer) SendMessage(ctx context.Context, key, value []byte) error {
	return p.kwriter.WriteMessages(ctx, kafka.Message{
		Key:   key,
		Value: value,
	})
}

// Close closes the producer writer.
func (p *KafkaProducer) Close() error {
	return p.kwriter.Close()
}

// // Connect to the specified topic and partition in the server
// func (w *KafkaProducer) Connect(topic string, partition int) (*kafka.Conn, error) {
// 	conn, err := kafka.DialLeader(context.Background(), "tcp",
// 		"localhost:9092", topic, partition)
// 	if err != nil {
// 		log.Println("failed to dial leader")
// 	}
// 	return conn, err
// }
