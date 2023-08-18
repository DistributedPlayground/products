package message

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/DistributedPlayground/go-lib/common"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
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
		return common.DPError(err)
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

	// Calculate the size of the message payload in bytes
	messageSize := len([]byte(valueJSON))
	fmt.Printf("Message size: %d bytes\n", messageSize)

	// Produce the message to the Kafka topic
	err = p.kp.Produce(&msg, nil)
	if err != nil {
		return common.DPError(err)
	}

	return nil
}
