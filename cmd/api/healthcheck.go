package main

import (
	"net/http"
)

func (app *application) healthCheckHandler(res http.ResponseWriter, req *http.Request) {

	//js := `{"status":"available","environment":%q,"version":%q}`
	//js = fmt.Sprintf(js, app.config.env, version)

	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}

	err := app.writeJSON(res, http.StatusOK, env, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(res, "Server encountered a problem, could not pass request", http.StatusInternalServerError)

	}

}
