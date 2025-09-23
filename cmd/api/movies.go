package main

import (
	"fmt"
	"net/http"
	"time"

	"movies.mlo_dev.net/internal/data"
)

func (app *application) createMovieHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "create a new movie")
}

func (app *application) showMovieHandler(res http.ResponseWriter, req *http.Request) {

	id, err := app.readIDParam(req)
	if err != nil || id < 1 {
		http.NotFound(res, req)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}

	err = app.writeJSON(res, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(res, "The server encountered a problem and could not request", http.StatusInternalServerError)
	}
}
