package store

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

var kp *kafka.Producer

func MustNewKafka() *kafka.Producer {
	if kp != nil {
		return kp
	}
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "kafka:9092"})
	if err != nil {
		panic(err)
	}
	kp = producer
	return kp
}
