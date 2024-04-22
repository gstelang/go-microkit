package main

// import (
// 	"database/sql"
// )

// main.go - entry point into that application
//  1. Initializes the microservice and starts as a separate executable.
//  2. Configures dependencies
//  3. Configures environment
//  3. Starts the server.

func main() {
	// Initialize a database connection
	// db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/database")
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()

	// Create a new instance of YourUserService
	// userRepositoryService := YourUserService(db)

	// Initialie a logger

	// Example using mux
	// mux.Handle("/health", handlers.HealthHandler(logger))
	// mux.Handle("GET /user/{userId}", handlers.CreateYourServiceHandler(userRepositoryService, logger))

	// // function options pattern to start the service.
	// // use a framework such as go-chi, etc
}
