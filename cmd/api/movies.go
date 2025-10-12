package main

import (
	"fmt"
	"net/http"
	"time"

	"movies.mlo_dev.net/internal/data"
	"movies.mlo_dev.net/internal/validator"
)

func (app *application) createMovieHandler(res http.ResponseWriter, req *http.Request) {
	//declare an anonymous struct to hold the info in http request body

	var input struct {
		Title   string       `json:"title"`
		Year    int32        `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genres"`
	}

	err := app.readJSON(res, req, &input)
	if err != nil {
		app.badRequestResponse(res, req, err)
		return
	}

	//Initialize a new Validator
	v := validator.New()

	//use the Check() method to execute our validation checks
	v.Check(input.Title != "", "title", "must be provided!")
	v.Check(len(input.Title) <= 500, "title", "must not be more than 500 bytes long")

	fmt.Fprintf(res, "%+v\n", input)

}

func (app *application) showMovieHandler(res http.ResponseWriter, req *http.Request) {

	id, err := app.readIDParam(req)
	if err != nil || id < 1 {
		//http.NotFound(res, req)
		app.notFoundResponse(res, req)
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
		//app.logger.Println(err)
		//http.Error(res, "The server encountered a problem and could not request", http.StatusInternalServerError)
		app.serverErrorResponse(res, req, err)
	}
}
