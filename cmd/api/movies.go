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

	movie := &data.Movie{
		Title:   input.Title,
		Year:    input.Year,
		Runtime: input.Runtime,
		Genres:  input.Genres,
	}

	//Initialize a new Validator
	v := validator.New()

	//use the Check() method to execute our validation checks
	v.Check(input.Title != "", "title", "must be provided!")
	v.Check(len(input.Title) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(input.Year != 0, "year", "must be provided")
	v.Check(input.Year >= 1888, "year", "must be greater than 1888")
	v.Check(input.Year <= int32(time.Now().Year()), "year", "must not be in the future")

	v.Check(input.Runtime != 0, "runtime", "must be provided")
	v.Check(input.Runtime > 0, "runtime", "must be a positive integer")

	v.Check(input.Genres != nil, "genres", "must be provided")
	v.Check(len(input.Genres) >= 1, "genres", "must contain at least 1 genre")
	v.Check(len(input.Genres) <= 5, "genres", "must not contain more than 5 genres")
	v.Check(validator.Unique(input.Genres), "genres", "must not contain duplicate values")

	//verify if any of the validation checks failed
	if !v.Valid() {
		app.failedValidationResponse(res, req, v.Errors)
		return
	}
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
