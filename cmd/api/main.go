package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// declare string containing application version number
const version = "1.0.0"

// define config struct to hold config settings
type config struct {
	port int
	env  string
}

// define struct to hold dependencies for http handler, helpers, and middleware
type application struct {
	config config
	logger *log.Logger
}

func main() {
	//config struct instance
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment(development|staging|production)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthCheckHandler)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	//start HTTP server
	logger.Printf("Starting %s server on port %d in %s mode", cfg.env, cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
