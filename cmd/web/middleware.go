package main

/*
import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

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
*/
