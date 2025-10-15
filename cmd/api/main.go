package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

// declare string containing application version number
const version = "1.0.0"

// define config struct to hold config settings
type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
}

// define struct to hold dependencies for http handler, helpers, and middleware
type application struct {
	config config
	logger *log.Logger
}

func main() {

	var cfg config

	dsn := os.Getenv("DSN")

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

	//Read the DSN from the db-dsn command-line flag into the config struct
	flag.StringVar(&cfg.db.dsn, "db-dsn", dsn, "PostgreSQL DSN")

	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	//call openDB() helper function to create connection pool
	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}

	//defer a call to db.Close() so connection pool is closed before
	//main() function exits.
	defer db.Close()

	//log a message to determine succesful establishment of connection pool
	logger.Printf("database connection pool established")

	app := &application{
		config: cfg,
		logger: logger,
	}

	//mux := http.NewServeMux()
	//mux.HandleFunc("/v1/healthcheck", app.healthCheckHandler)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	//start HTTP server
	logger.Printf("Starting %s server on port %d", cfg.env, cfg.port)
	err = srv.ListenAndServe()
	logger.Fatal(err)
}

func openDB(cfg config) (*sql.DB, error) {
	//use sql.Open() to create an empty connection pool
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	//create a context with S-second timeout deadline.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//use PingContext() to establish a new connection to the database
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	//return the sql.DB connection pool
	return db, nil
}

//find . -name "*.go" | entr -r sh -c 'echo "== Restarting =="; go run ./cmd/api -db-dsn 'user=greenlight password=pa$$word host=localhost dbname=greenlight sslmode=disable''
