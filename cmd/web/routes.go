package main

import (
	"ecommerce/docs"
	"fmt"
	"net/http"
	"time"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func (app *application) route() *http.ServeMux {
	router := http.NewServeMux()

	// Product routes
	router.Handle("GET /product/{prod_id}", app.prodContextMiddleware(http.HandlerFunc(app.getProduct)))
	router.Handle("DELETE /product/{prod_id}", app.prodContextMiddleware(http.HandlerFunc(app.deleteProduct)))
	router.Handle("PATCH /product/{prod_id}", app.prodContextMiddleware(http.HandlerFunc(app.updateProduct)))
	router.Handle("POST /product", http.HandlerFunc(app.addProduct))

	// User routes
	router.Handle("GET /user/{user_id}", app.userContextMiddleware(http.HandlerFunc(app.getUser)))
	router.Handle("DELETE /user/{user_id}", app.userContextMiddleware(http.HandlerFunc(app.deleteUser)))
	router.Handle("PATCH /user/{user_id}", app.userContextMiddleware(http.HandlerFunc(app.updateUser)))
	router.Handle("POST /user", http.HandlerFunc(app.addUser))

	// Swagger docs
	docsURL := fmt.Sprintf("%s/swagger/doc.json", app.config.addr)
	router.Handle("/swagger/*", httpSwagger.Handler(httpSwagger.URL(docsURL)))

	v1 := http.NewServeMux()
	v1.Handle("/v1/", http.StripPrefix("/v1", router))

	return v1
}

func (app *application) run(mux http.Handler) error {
	// Docs metadata
	docs.SwaggerInfo.Title = "Ecommerce API"
	docs.SwaggerInfo.Version = version
	docs.SwaggerInfo.Host = app.config.apiURL
	docs.SwaggerInfo.BasePath = "/v1"

	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	return srv.ListenAndServe()
}
