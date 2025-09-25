package store

import (
	"context"
	"database/sql"
	"time"
)

type Order struct {
	ID        int
	UserID    int
	Total     float64
	CreatedAt time.Time
	Items     []OrderItem
}

type OrderStore struct {
	db *sql.DB
}

func (o *OrderStore) Create(ctx context.Context, order *Order) error {
	return nil
}
