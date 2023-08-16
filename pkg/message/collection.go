package message

import (
	"encoding/json"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/segmentio/ksuid"
)

type Collection interface {
	Send(value interface{}, messageType string) error
}

type collection struct {
	kp    *kafka.Producer
	topic string
}

func NewCollection(kp *kafka.Producer) Collection {
	return &collection{kp: kp, topic: "collection"}
}

func (c collection) Send(value interface{}, messageType string) error {
	// Generate a UUID for the message order identifier
	messageOrderID := ksuid.New().String()

	valueJSON, err := json.Marshal(value)
	if err != nil {
		return err
	}

	key := []byte(messageOrderID) // Use the UUID as the message key

	msg := kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &c.topic, Partition: kafka.PartitionAny},
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
	err = c.kp.Produce(&msg, nil)
	if err != nil {
		return err
	}

	return nil
}
