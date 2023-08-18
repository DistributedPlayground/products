package service

import (
	"context"

	"github.com/DistributedPlayground/go-lib/common"
	"github.com/DistributedPlayground/products/pkg/message"
	"github.com/DistributedPlayground/products/pkg/model"
	"github.com/DistributedPlayground/products/pkg/repository"
)

type Collection interface {
	Create(ctx context.Context, request model.CollectionUpsert) (model.Collection, error)
	Update(ctx context.Context, id string, updates model.CollectionUpsert) (model.Collection, error)
}

type collection struct {
	repo    repository.Collection
	message message.Collection
}

func NewCollection(repo repository.Collection, message message.Collection) Collection {
	return &collection{repo: repo, message: message}
}

func (c collection) Create(ctx context.Context, request model.CollectionUpsert) (model.Collection, error) {
	collection, err := c.repo.Create(ctx, request)
	if err != nil {
		return model.Collection{}, common.DPError(err)
	}

	// Send to Kafka
	err = c.message.Send(collection, "Create")
	if err != nil {
		return collection, common.DPError(err)
	}

	return collection, nil
}

func (c collection) Update(ctx context.Context, id string, updates model.CollectionUpsert) (model.Collection, error) {
	collection, err := c.repo.Update(ctx, id, updates)
	if err != nil {
		return collection, common.DPError(err)
	}

	// Send to Kafka
	err = c.message.Send(collection, "Update")
	if err != nil {
		return collection, common.DPError(err)
	}

	return collection, nil
}
