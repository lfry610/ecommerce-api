package main

import (
	"net/http"
)

func (app *application) InternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Errorw("internal error", "method", r.Method, "path", r.URL.Path, "error", err)

	writeJSONError(w, http.StatusInternalServerError, "the server encountered a problem")
}

func (app *application) StatusBadRequest(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Errorw("bad request error", "method", r.Method, "path", r.URL.Path, "error", err)

	writeJSONError(w, http.StatusBadRequest, "the server encountered a problem")
}

func (app *application) StatusNotFound(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Errorw("not found error", "method", r.Method, "path", r.URL.Path, "error", err)

	writeJSONError(w, http.StatusNotFound, "the server encountered a problem")
}
