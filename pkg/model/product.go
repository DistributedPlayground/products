package model

import (
	"time"
)

type Product struct {
	Id           string     `json:"id" db:"id"`
	CreatedAt    time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt    time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt    *time.Time `json:"deletedAt,omitempty" db:"deleted_at"`
	Name         string     `json:"name" db:"name"`
	Description  string     `json:"description" db:"description"`
	Price        float64    `json:"price" db:"price"`
	Inventory    int        `json:"inventory" db:"inventory"`
	CollectionId *string    `json:"collectionId,omitempty" db:"collection_id"`
}

type ProductUpsert struct {
	Name         *string  `json:"name" db:"name"`
	Description  *string  `json:"description" db:"description"`
	Price        *float64 `json:"price" db:"price"`
	Inventory    *int     `json:"inventory" db:"inventory"`
	CollectionId *string  `json:"collectionId" db:"collection_id"`
}
