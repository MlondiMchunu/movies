package main

import "net/http"

func (app *application) logError(r *http.Request, err error) {
	app.logger.Println(err)
}

func (app *application) errorResponse(res http.ResponseWriter, req *http.Request, status int, message interface{}) {
	env := envelope{"error": message}

	err := app.writeJSON(res, status, env, nil)
	if err != nil {
		app.logError(req, err)
		res.WriteHeader(500)
	}
}
