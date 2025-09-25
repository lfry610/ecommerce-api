package store

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"time"
)

type User struct {
	ID           int       `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
}

type UserStore struct {
	db *sql.DB
}

func (u *UserStore) Create(ctx context.Context, user *User) error {
	query := `
		INSERT INTO users (email, password_hash)
			VALUES ($1, $2) RETURNING id, created_at
		`

	err := u.db.QueryRowContext(
		ctx,
		query,
		user.Email,
		user.PasswordHash,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	)

	if err != nil {
		return err
	}
	return nil
}

func (u *UserStore) GetUser(ctx context.Context, id int) (*User, error) {
	query := `
		SELECT id, email, password_hash, created_at
			FROM users
				WHERE id = $1
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	var user User
	err := u.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, http.ErrNotSupported
		default:
			return nil, err
		}
	}
	return &user,
		nil
}

func (u *UserStore) Delete(ctx context.Context, id int) error {
	query := `
	DELETE FROM users WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	res, err := u.db.ExecContext(ctx, query, id)
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

func (u *UserStore) Update(ctx context.Context, user *User) error {
	query := `
		UPDATE users
		SET email = $1, password_hash = $2
		WHERE id = $3
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := u.db.ExecContext(ctx, query, user.Email, user.PasswordHash, user.ID)
	if err != nil {
		return err
	}

	return nil
}
