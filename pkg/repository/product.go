package repository

import (
	"context"

	"github.com/DistributedPlayground/go-lib/database"

	"github.com/DistributedPlayground/products/pkg/model"
)

type Product interface {
	database.Transactable
	Create(ctx context.Context, product model.ProductUpsert) (model.Product, error)
	Update(ctx context.Context, id string, updates any) (model.Product, error)
}

type product[T any] struct {
	Base[T]
}

func NewProduct(db database.Queryable) Product {
	return &product[model.Product]{Base[model.Product]{Store: db, Table: "product"}}
}

func (c product[T]) Create(ctx context.Context, request model.ProductUpsert) (product model.Product, err error) {
	query, args, err := c.Named(`
		INSERT INTO product (name, description, inventory, price, collection_id)
		VALUES (:name, :description, :inventory, :price, :collection_id) RETURNING *`, request)

	if err != nil {
		return model.Product{}, err
	}

	err = c.Store.QueryRowxContext(ctx, query, args...).StructScan(&product)
	if err != nil {
		return model.Product{}, err
	}

	return product, nil
}
