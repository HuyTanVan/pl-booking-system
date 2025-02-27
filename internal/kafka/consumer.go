package kafka_service

import (
	"context"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	kreader *kafka.Reader
}

// NewConsumer creates a new consumer instance.
func NewConsumer(cfg Config, cCfg ConsumerConfig) *KafkaConsumer {
	return &KafkaConsumer{
		kreader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:        cfg.Brokers,
			Topic:          cfg.Topic,
			GroupID:        cCfg.GroupID,
			Partition:      cCfg.Partition,
			MinBytes:       cCfg.MinBytes,
			MaxBytes:       cCfg.MaxBytes,
			CommitInterval: time.Duration(cCfg.CommitInterval) * time.Millisecond,
			StartOffset:    cCfg.StartOffset,
		}),
	}
}

// ReadMessage reads a message from the Kafka topic.
func (c *KafkaConsumer) ReadMessage(ctx context.Context) (kafka.Message, error) {
	return c.kreader.ReadMessage(ctx)
}

// Close closes the consumer reader.
func (c *KafkaConsumer) Close() error {
	return c.kreader.Close()
}
