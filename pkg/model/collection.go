package model

import "time"

type Collection struct {
	Id          string     `json:"id" db:"id"`
	CreatedAt   time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt   *time.Time `json:"deletedAt,omitempty" db:"deleted_at"`
	Name        string     `json:"name" db:"name"`
	Description string     `json:"description" db:"description"`
}

type CollectionUpdates struct {
	Name        *string `json:"name" db:"name"`
	Description *string `json:"description" db:"description"`
}
