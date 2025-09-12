package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) healthCheckHandler(res http.ResponseWriter, req *http.Request) {

	//js := `{"status":"available","environment":%q,"version":%q}`
	//js = fmt.Sprintf(js, app.config.env, version)

	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	js, err := json.Marshal(data)
	if err != nil {
		app.logger.Println(err)
		http.Error(res, "Server encounteed a problem, could not pass request", http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")

	res.Write([]byte(js))
}
