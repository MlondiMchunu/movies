package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthCheckHandler(res http.ResponseWriter, req *http.Request) {
	js := `{"status":"available","environment":%q,"version":%q}`
	js = fmt.Sprintf(js, app.config.env, version)

	res.Header().Set("Content-Type", "application/json")

	res.Write([]byte(js))
}
