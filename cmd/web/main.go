package main

import (
	"ecommerce/internal/db"
	"ecommerce/internal/env"
	"ecommerce/internal/store"

	"go.uber.org/zap"
)

type application struct {
	logger *zap.SugaredLogger
	config config
	store  store.Storage
}

type config struct {
	addr   string
	db     dbConfig
	env    string
	apiURL string
}
type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

//	@title			Ecommerce API
//	@description	API for an ECommerce shop
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@BasePath					/v1
//
//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description

func main() {
	cfg := config{
		addr:   env.GetString("addr", ":4000"),
		apiURL: env.GetString("EXTERNAL_URL", "localhost:4000"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/admin_db?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env: env.GetString("ENV", "development"),
	}

	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)

	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()
	logger.Info("database has connected")

	store := store.NewPostgresStorage(db)
	app := &application{
		logger: logger,
		config: cfg,
		store:  store,
	}

	logger.Infow("server has started", "addr", app.config.addr, "env", app.config.env)
	mux := app.route()

	logger.Fatal(app.run(mux))
}
