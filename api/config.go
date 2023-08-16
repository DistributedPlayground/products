package api

import (
	"github.com/DistributedPlayground/products/pkg/message"
	"github.com/DistributedPlayground/products/pkg/repository"
	"github.com/DistributedPlayground/products/pkg/service"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jmoiron/sqlx"
)

func NewRepos(db *sqlx.DB) repository.Repositories {
	return repository.Repositories{
		Collection: repository.NewCollection(db),
		Product:    repository.NewProduct(db),
	}
}

func NewServices(repos repository.Repositories, messages message.Messages) service.Services {
	return service.Services{
		Collection: service.NewCollection(repos.Collection, messages.Collection),
		Product:    service.NewProduct(repos.Product, messages.Product),
	}
}

func NewMessages(kp *kafka.Producer) message.Messages {
	return message.Messages{
		Collection: message.NewCollection(kp),
		Product:    message.NewProduct(kp),
	}
}
