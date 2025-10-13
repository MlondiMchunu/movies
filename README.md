# Movies API

A RESTful API for managing a collection of movies, built with Go.

## Features
- CRUD operations for movies
- Healthcheck endpoint
- Input validation
- Error handling
- Modular project structure

## Project Structure
```
go.mod
Makefile
README.md
bin/
cmd/
  api/
    errors.go
    healthcheck.go
    helpers.go
    main.go
    movies.go
    routes.go
internal/
  data/
    movies.go
    runtime.go
  validator/
    validator.go
migrations/
remote/
```

## Getting Started

### Prerequisites
- Go 1.18+

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/MlondiMchunu/movies.git
   cd movies
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```

### Running the API
1. Build the project:
   ```bash
   make build
   ```
2. Run the API server:
   ```bash
   make run
   ```
   or
   ```bash
   go run ./cmd/api/main.go
   ```

### API Endpoints
- `GET /v1/healthcheck` — Healthcheck endpoint
- `GET /v1/movies` — List all movies
- `POST /v1/movies` — Create a new movie
- `GET /v1/movies/{id}` — Get a movie by ID
- `PUT /v1/movies/{id}` — Update a movie
- `DELETE /v1/movies/{id}` — Delete a movie

## Development
- Code is organized into `cmd/api` for the main API logic and `internal` for supporting packages.
- Migrations are stored in the `migrations/` directory.
- Use the Makefile for common tasks.

## License
MIT

