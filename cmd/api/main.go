package main

import (
	"flag"
	"log"
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
}
