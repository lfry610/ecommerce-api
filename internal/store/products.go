package store

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"time"
)

type Product struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Stock     int       `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
}

type ProductStore struct {
	db *sql.DB
}

func (p *ProductStore) Create(ctx context.Context, product *Product) error {
	query := `
		INSERT INTO products (id, name, price, stock)
			VALUES ($1, $2, $3, $4) RETURNING created_at
		`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := p.db.QueryRowContext(
		ctx,
		query,
		product.ID,
		product.Name,
		product.Price,
		product.Stock,
	).Scan(
		&product.CreatedAt,
	)

	if err != nil {
		return err
	}
	return nil
}

func (p *ProductStore) GetProdByID(ctx context.Context, id int) (*Product, error) {
	query := `
		SELECT id, name, price, stock, created_at
			FROM products
				WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	var product Product
	err := p.db.QueryRowContext(ctx, query, id).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.Stock,
		&product.CreatedAt,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, http.ErrNotSupported
		default:
			return nil, err
		}
	}
	return &product,
		nil
}

func (p *ProductStore) Delete(ctx context.Context, id int) error {
	query := `
	DELETE FROM products WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	res, err := p.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrNotFound
	}

	return nil

}

func (p *ProductStore) Update(ctx context.Context, product *Product) error {
	query := `
		UPDATE products
		SET name = $1, price = $2, stock = $3
		WHERE id = $4
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := p.db.ExecContext(ctx, query, product.Name, product.Price, product.Stock, product.ID)
	if err != nil {
		return err
	}

	return nil
}
