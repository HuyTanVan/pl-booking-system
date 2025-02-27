package kafka_service

// Config common configuration used by both producers and consumers.
type Config struct {
	Brokers []string // list of server addresses: ["kafka:9092", "kafka:9093"]
	Topic   string
}

// ProducerConfig specific configuration for producers.
type ProducerConfig struct {
	BatchSize    int // numbers of messages to collect before sending 'em to kafka
	BatchTimeout int // In milliseconds
	Async        bool
	RequiredAcks int //Specifies the number of broker acknowledgements required before considering a message as sent
}

// ConsumerConfig specific configuration for consumers.
type ConsumerConfig struct {
	GroupID        string
	Partition      int
	MinBytes       int
	MaxBytes       int
	CommitInterval int // In milliseconds
	StartOffset    int64
}
