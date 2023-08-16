package message

import (
	"encoding/json"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/segmentio/ksuid"
)

type Product interface {
	Send(value interface{}, messageType string) error
}

type product struct {
	kp    *kafka.Producer
	topic string
}

func NewProduct(kp *kafka.Producer) Product {
	return &product{kp: kp, topic: "product"}
}

func (p product) Send(value interface{}, messageType string) error {
	// Generate a UUID for the message order identifier
	messageOrderID := ksuid.New().String()

	key := []byte(messageOrderID) // Use the UUID as the message key

	valueJSON, err := json.Marshal(value)
	if err != nil {
		return err
	}

	msg := kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &p.topic, Partition: kafka.PartitionAny},
		Key:            []byte(key),
		Value:          valueJSON,
		Headers:        nil,
		Timestamp:      time.Time{},
	}

	// Include the message type in the headers or value
	msg.Headers = append(msg.Headers, kafka.Header{
		Key:   "MessageType",
		Value: []byte(messageType),
	})

	// Produce the message to the Kafka topic
	err = p.kp.Produce(&msg, nil)
	if err != nil {
		return err
	}

	return nil
}
