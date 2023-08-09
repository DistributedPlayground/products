package repository

import (
	"context"

	"github.com/DistributedPlayground/go-lib/database"

	"github.com/DistributedPlayground/products/pkg/model"
)

type Collection interface {
	database.Transactable
	Create(ctx context.Context, collection model.Collection) (model.Collection, error)
	Update(ctx context.Context, id string, updates any) error
}

type collection[T any] struct {
	Base[T]
}

func NewCollection(db database.Queryable) Collection {
	return &collection[model.Collection]{Base[model.Collection]{Store: db, Table: "collection"}}
}

func (c collection[T]) Create(ctx context.Context, collection model.Collection) (model.Collection, error) {
	query, args, err := c.Named(`
		INSERT INTO collection (name, description)
		VALUES (:name, :description) RETURNING *`, collection)

	if err != nil {
		return model.Collection{}, err
	}

	err = c.Store.QueryRowxContext(ctx, query, args...).StructScan(&collection)
	if err != nil {
		return model.Collection{}, err
	}

	return collection, nil
}
