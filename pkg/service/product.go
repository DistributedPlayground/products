package service

import (
	"context"

	"github.com/DistributedPlayground/products/pkg/model"
	"github.com/DistributedPlayground/products/pkg/repository"
)

type Product interface {
}

type product struct {
	repo repository.Product
}

func NewProduct(repo repository.Product) Product {
	return &product{repo: repo}
}

func (c product) Create(ctx context.Context, request model.Product) (model.Product, error) {
	product, err := c.repo.Create(ctx, request)
	if err != nil {
		return model.Product{}, err
	}
	// TODO: Add event to queue
	return product, nil
}

func (c product) Update(ctx context.Context, id string, updates model.ProductUpdates) error {
	err := c.repo.Update(ctx, id, updates)
	if err != nil {
		return err
	}
	// TODO: Add event to queue
	return nil
}