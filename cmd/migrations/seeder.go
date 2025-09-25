package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/lib/pq"
)

// --- Product model & store (simplified from your code) ---
type Product struct {
	ID        int
	Name      string
	Price     float64
	Stock     int
	CreatedAt time.Time
}

type ProductStore struct {
	db *sql.DB
}

func (p *ProductStore) Create(ctx context.Context, product *Product) error {
	query := `
		INSERT INTO products (id, name, price, stock)
		VALUES ($1, $2, $3, $4)
		RETURNING created_at
	`
	return p.db.QueryRowContext(
		ctx,
		query,
		product.ID,
		product.Name,
		product.Price,
		product.Stock,
	).Scan(&product.CreatedAt)
}

// --- Seeder ---
func main() {
	// ⚡ Update with your DB credentials
	dsn := "postgres://admin:adminpassword@localhost/admin_db?sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("failed to connect to db:", err)
	}
	defer db.Close()

	store := &ProductStore{db: db}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Fake product names
	names := []string{
		"Wireless Mouse",
		"Mechanical Keyboard",
		"Gaming Headset",
		"4K Monitor",
		"USB-C Charger",
		"Laptop Stand",
		"Portable SSD",
		"Smartphone Case",
		"Bluetooth Speaker",
		"Webcam HD",
		"PS5",
		"Nintendo Switch",
		"Blu Ray Player",
		"Xbox Series S",
		"Macbook Air M3",
	}

	// Insert N fake products
	for i := 0; i < len(names); i++ {
		product := &Product{
			ID:    rand.Intn(100000),               // random ID
			Name:  names[i],                        // product name
			Price: float64(rand.Intn(20000)) / 100, // price between 0–200
			Stock: rand.Intn(100) + 1,              // stock 1–100
		}

		if err := store.Create(ctx, product); err != nil {
			log.Println("Insert failed:", err)
			continue
		}
		log.Printf("Inserted product: %s ($%.2f, stock=%d, id=%d)", product.Name, product.Price, product.Stock, product.ID)
	}

	fmt.Println("Seeding complete 🎉")
}
