package main

import (
	"net/http"

	"github.com/Ruthvik10/go-shared-library/jsonutil"
)

type envelope map[string]any

func (app *application) logError(r *http.Request, err error) {
	app.logger.Error(err, map[string]any{
		"request_method": r.Method,
		"request_url":    r.URL.String(),
	})
}

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	err := jsonutil.WriteJSON(w, envelope{"error": message}, status, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *application) ServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)
	message := "the server encountered a problem and could not process your request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

func (app *application) NotFoundErrorResponse(w http.ResponseWriter, r *http.Request) {
	message := "could not find the requested resource"
	app.errorResponse(w, r, http.StatusNotFound, message)
}

func (app *application) BadRequestErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.errorResponse(w, r, http.StatusNotFound, err.Error())
}
