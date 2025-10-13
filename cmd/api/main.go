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

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// declare string containing application version number
const version = "1.0.0"

// define config struct to hold config settings
type config struct {
	port int
	env  string
	db   struct {
		user     string
		password string
		host     string
		name     string
		dsn      string
	}
}

// define struct to hold dependencies for http handler, helpers, and middleware
type application struct {
	config config
	logger *log.Logger
}

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found or could not be loaded")
	}

	var cfg config

	// Read environment variables
	cfg.port = 4000
	if portEnv := os.Getenv("PORT"); portEnv != "" {
		fmt.Sscanf(portEnv, "%d", &cfg.port)
	}
	cfg.env = os.Getenv("ENV")
	if cfg.env == "" {
		cfg.env = "development"
	}
	cfg.db.user = os.Getenv("DB_USER")
	cfg.db.password = os.Getenv("DB_PASSWORD")
	cfg.db.host = os.Getenv("DB_HOST")
	cfg.db.name = os.Getenv("DB_NAME")

	// Build DSN from env variables if present
	if cfg.db.user != "" && cfg.db.password != "" && cfg.db.host != "" && cfg.db.name != "" {
		cfg.db.dsn = fmt.Sprintf("postgres://%s:%s@%s/%s", cfg.db.user, cfg.db.password, cfg.db.host, cfg.db.name)
	} else if dsn := os.Getenv("DB_DSN"); dsn != "" {
		cfg.db.dsn = dsn
	} else {
		cfg.db.dsn = "postgres://greenlight:pa$$word@localhost/greenlight"
	}

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

//find . -name "*.go" | entr -r sh -c 'echo "== Restarting =="; go run ./cmd/api'
