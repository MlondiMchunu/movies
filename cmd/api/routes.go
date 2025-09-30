package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	//initialize a new httprouter router instance
	router := httprouter.New()

	//convert the notFoundResponse() helper to a http.Handler
	//using http.HandlerFunc() adapter
	router.NotFound = http.HandlerFunc(app.notFoundResponse)

	//register relevant methods, URL patterns and handler functions
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthCheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)

	return router
}
