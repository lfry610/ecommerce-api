package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrNotFound          = errors.New("resource not found")
	QueryTimeoutDuration = time.Second * 5
)

type Storage struct {
	OrderItem interface {
		Create(context.Context, *OrderItem) error
	}
	Product interface {
		Create(context.Context, *Product) error
		GetProdByID(context.Context, int) (*Product, error)
		Delete(context.Context, int) error
		Update(context.Context, *Product) error
	}
	Users interface {
		Create(context.Context, *User) error
		GetUser(context.Context, int) (*User, error)
		Delete(context.Context, int) error
		Update(context.Context, *User) error
	}
	Order interface {
		Create(context.Context, *Order) error
	}
}

func NewPostgresStorage(db *sql.DB) Storage {
	return Storage{
		OrderItem: &OrderItemStore{db},
		Product:   &ProductStore{db},
		Users:     &UserStore{db},
		Order:     &OrderStore{db},
	}
}
