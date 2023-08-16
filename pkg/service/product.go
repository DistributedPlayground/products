package service

import (
	"context"

	"github.com/DistributedPlayground/products/pkg/message"
	"github.com/DistributedPlayground/products/pkg/model"
	"github.com/DistributedPlayground/products/pkg/repository"
)

type Product interface {
	Create(ctx context.Context, request model.ProductUpsert) (model.Product, error)
	Update(ctx context.Context, id string, updates model.ProductUpsert) (model.Product, error)
}

type product struct {
	repo    repository.Product
	message message.Product
}

func NewProduct(repo repository.Product, message message.Product) Product {
	return &product{repo, message}
}

func (c product) Create(ctx context.Context, request model.ProductUpsert) (model.Product, error) {
	product, err := c.repo.Create(ctx, request)
	if err != nil {
		return model.Product{}, err
	}

	// Send to Kafka
	err = c.message.Send(product, "Create")
	if err != nil {
		return product, err
	}

	return product, nil
}

func (c product) Update(ctx context.Context, id string, updates model.ProductUpsert) (model.Product, error) {
	product, err := c.repo.Update(ctx, id, updates)
	if err != nil {
		return product, err
	}

	// Send to Kafka
	err = c.message.Send(product, "Update")
	if err != nil {
		return product, err
	}

	return product, nil
}
