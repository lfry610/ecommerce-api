package store

import (
	"context"
	"database/sql"
)

type OrderItem struct {
	ID        int
	OrderID   int
	ProductID int
	Qty       int
	Price     float64
}

type OrderItemStore struct {
	db *sql.DB
}

func (o *OrderItemStore) Create(ctx context.Context, orderItem *OrderItem) error {
	return nil
	/*
		query := `
			INSERT INTO OrderItem (ID, OrderID, ProductID, Qty, Price)
				VALUES ($1, $2, $3, $4, $5)
		`
	*/
}
