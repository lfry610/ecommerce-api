package main

import (
	"log"
	"net/http"
)

func (app *application) InternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("internal server error: %s path: %s error: %s", r.Method, r.URL.Path, err)

	writeJSONError(w, http.StatusInternalServerError, "the server encountered a problem")
}

func (app *application) StatusBadRequest(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("internal server error: %s path: %s error: %s", r.Method, r.URL.Path, err)

	writeJSONError(w, http.StatusBadRequest, "the server encountered a problem")
}

func (app *application) StatusNotFound(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("internal server error: %s path: %s error: %s", r.Method, r.URL.Path, err)

	writeJSONError(w, http.StatusNotFound, "the server encountered a problem")
}
