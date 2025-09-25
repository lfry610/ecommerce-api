package main

import (
	"ecommerce/internal/db"
	"ecommerce/internal/env"
	"ecommerce/internal/store"
	"log"
	"log/slog"
	"os"
)

type application struct {
	logger *slog.Logger
	config config
	store  store.Storage
}

type config struct {
	addr string
	db   dbConfig
	env  string
}
type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

func main() {
	cfg := config{
		addr: env.GetString("addr", ":4000"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/admin_db?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env: env.GetString("ENV", "development"),
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)

	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Println("database connected")

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	store := store.NewPostgresStorage(db)
	app := &application{
		logger: logger,
		config: cfg,
		store:  store,
	}

	logger.Info("Serving on port", slog.Any("addr", ":4000"))
	mux := app.route()

	log.Fatal(app.run(mux))
	logger.Error(err.Error())
	os.Exit(1)
}
