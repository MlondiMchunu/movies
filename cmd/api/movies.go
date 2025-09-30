package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"movies.mlo_dev.net/internal/data"
)

func (app *application) createMovieHandler(res http.ResponseWriter, req *http.Request) {
	//declare an anonymous struct to hold the info in http request body

	var input struct {
		Title   string   `json:"title"`
		Year    int32    `json:"year"`
		Runtime int32    `json:"runtime"`
		Genres  []string `json:"genres"`
	}

	//initialize new json.Decoder instance which reads from request body

	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		app.errorResponse(res, req, http.StatusBadRequest, err.Error())
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
