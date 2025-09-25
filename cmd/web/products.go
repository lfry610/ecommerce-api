package main

import (
	"context"
	"ecommerce/internal/store"
	"errors"
	"math/rand/v2"
	"net/http"
	"strconv"
)

type AddProductPayload struct {
	Name  string  `json:"name" validate:"required"`
	Price float64 `json:"price" validate:"required"`
	Stock int     `json:"stock" validate:"required"`
}

func (app *application) getProduct(w http.ResponseWriter, r *http.Request) {
	product := getProdFromCtx(r)

	if err := app.jsonResponse(w, http.StatusOK, product); err != nil {
		app.InternalServerError(w, r, err)
		return
	}
}

func (app *application) addProduct(w http.ResponseWriter, r *http.Request) {
	var payload AddProductPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.StatusBadRequest(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.StatusBadRequest(w, r, err)
		return
	}

	product := &store.Product{
		ID:    rand.IntN(10000),
		Name:  payload.Name,
		Price: payload.Price,
		Stock: payload.Stock,
	}

	ctx := r.Context()

	if err := app.store.Product.Create(ctx, product); err != nil {
		app.InternalServerError(w, r, err)
	}

	if err := app.jsonResponse(w, http.StatusCreated, product); err != nil {
		app.InternalServerError(w, r, err)
	}
}

func (app *application) deleteProduct(w http.ResponseWriter, r *http.Request) {

	prodID, err := strconv.Atoi(r.PathValue("prod_id"))
	if err != nil || prodID < 1 {
		app.InternalServerError(w, r, err)
		return
	}
	ctx := r.Context()

	if err := app.store.Product.Delete(ctx, prodID); err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.StatusNotFound(w, r, err)
		default:
			app.InternalServerError(w, r, err)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

type UpdateProdPayload struct {
	Name  string  `json:"name" validate:"required,max=100"`
	Price float64 `json:"price" validate:"required,max=1000"`
	Stock int     `json:"stock" validate:"required,max=100"`
}

func (app *application) updateProduct(w http.ResponseWriter, r *http.Request) {
	product := getProdFromCtx(r)

	var payload UpdateProdPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.StatusBadRequest(w, r, err)
	}

	if err := Validate.Struct(payload); err != nil {
		app.StatusBadRequest(w, r, err)
	}

	product.Name = payload.Name
	product.Price = payload.Price
	product.Stock = payload.Stock

	err := app.store.Product.Update(r.Context(), product)
	if err != nil {
		app.InternalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, product); err != nil {
		app.InternalServerError(w, r, err)
	}

}

func (app *application) prodContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prodID, err := strconv.Atoi(r.PathValue("prod_id"))
		if err != nil || prodID < 1 {
			app.InternalServerError(w, r, err)
			return
		}

		ctx := r.Context()

		product, err := app.store.Product.GetProdByID(ctx, prodID)
		if err != nil {
			switch {
			case errors.Is(err, http.ErrAbortHandler):
				app.StatusNotFound(w, r, err)
			default:
				app.InternalServerError(w, r, err)
			}
			return
		}

		ctx = context.WithValue(ctx, "product", product)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getProdFromCtx(r *http.Request) *store.Product {
	product, _ := r.Context().Value("product").(*store.Product)
	return product
}
