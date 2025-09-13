package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// retrieve "id" URL parameter from current req context, then convert it to int
func (app *application) readIDParam(req *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(req.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}
	return id, nil
}

func (app *application) witeJSON(res http.ResponseWriter, status int, data interface{}, headers http.Header) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	//append new line to make it easier to view in terminal applications
	js = append(js, '\n')
}
