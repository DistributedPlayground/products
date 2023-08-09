package service

import (
	"context"

	"github.com/DistributedPlayground/products/pkg/model"
	"github.com/DistributedPlayground/products/pkg/repository"
)

type Collection interface {
	Create(ctx context.Context, request model.CollectionUpsert) (model.Collection, error)
	Update(ctx context.Context, id string, updates model.CollectionUpsert) error
}

type collection struct {
	repo repository.Collection
}

func NewCollection(repo repository.Collection) Collection {
	return &collection{repo: repo}
}

func (c collection) Create(ctx context.Context, request model.CollectionUpsert) (model.Collection, error) {
	collection, err := c.repo.Create(ctx, request)
	if err != nil {
		return model.Collection{}, err
	}
	// TODO: Add event to queue
	return collection, nil
}

func (c collection) Update(ctx context.Context, id string, updates model.CollectionUpsert) error {
	err := c.repo.Update(ctx, id, updates)
	if err != nil {
		return err
	}
	// TODO: Add event to queue
	return nil
}
