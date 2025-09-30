package main

import (
	"context"
	"ecommerce/internal/store"
	"errors"
	"math/rand/v2"
	"net/http"
	"strconv"
)

const version = "1.0.0"

type AddUserPayload struct {
	Email        string `json:"email" validate:"required"`
	PasswordHash string `json:"password_hash" validate:"required"`
}

func (app *application) getUser(w http.ResponseWriter, r *http.Request) {
	user := getUserFromCtx(r)

	if err := app.jsonResponse(w, http.StatusOK, user); err != nil {
		app.InternalServerError(w, r, err)
		return
	}
}

func (app *application) addUser(w http.ResponseWriter, r *http.Request) {
	var payload AddUserPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.StatusBadRequest(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.StatusBadRequest(w, r, err)
		return
	}

	user := &store.User{
		ID:           rand.IntN(10000),
		Email:        payload.Email,
		PasswordHash: payload.PasswordHash,
	}

	ctx := r.Context()

	if err := app.store.Users.Create(ctx, user); err != nil {
		app.InternalServerError(w, r, err)
	}

	if err := app.jsonResponse(w, http.StatusCreated, user); err != nil {
		app.InternalServerError(w, r, err)
	}
}

func (app *application) deleteUser(w http.ResponseWriter, r *http.Request) {

	userID, err := strconv.Atoi(r.PathValue("user_id"))
	if err != nil || userID < 1 {
		app.InternalServerError(w, r, err)
		return
	}
	ctx := r.Context()

	if err := app.store.Users.Delete(ctx, userID); err != nil {
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

type UpdateUserPayload struct {
	Email        string `json:"email" validate:"required"`
	PasswordHash string `json:"password_hash" validate:"required"`
}

func (app *application) updateUser(w http.ResponseWriter, r *http.Request) {
	user := getUserFromCtx(r)

	var payload UpdateUserPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.StatusBadRequest(w, r, err)
	}

	if err := Validate.Struct(payload); err != nil {
		app.StatusBadRequest(w, r, err)
	}

	user.Email = payload.Email
	user.PasswordHash = payload.PasswordHash

	err := app.store.Users.Update(r.Context(), user)
	if err != nil {
		app.InternalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, user); err != nil {
		app.InternalServerError(w, r, err)
	}

}

func (app *application) userContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, err := strconv.Atoi(r.PathValue("user_id"))
		if err != nil || userID < 1 {
			app.InternalServerError(w, r, err)
			return
		}

		ctx := r.Context()

		user, err := app.store.Users.GetUser(ctx, userID)
		if err != nil {
			switch {
			case errors.Is(err, http.ErrAbortHandler):
				app.StatusNotFound(w, r, err)
			default:
				app.InternalServerError(w, r, err)
			}
			return
		}

		ctx = context.WithValue(ctx, "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUserFromCtx(r *http.Request) *store.User {
	user, _ := r.Context().Value("user").(*store.User)
	return user
}
