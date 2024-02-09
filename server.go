package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/dev-xero/authentication-backend/application"
	"github.com/dev-xero/authentication-backend/database"
)

/*
Connects to the PostGreSQL database using the credentials

Params:
  - No parameters

Returns:
  - A pointer to the SQL database
  - An error if the connection failed
*/
func initDatabaseConnection() (*sql.DB, error) {
	db, err := database.ConnectDatabase()
	if err != nil {
		msg := "[FAIL]: unable to connect database"
		return nil, fmt.Errorf("%s: %w", msg, err)
	}

	return db, nil
}

/*
Main is the server entry point

Objectives:
  - Initializes a database connection
  - Spins up a context channel to handle OS interrupts
  - Starts the server

Params:
  - No parameters

Returns:
  - No return value
*/
func main() {
	appDatabase, err := initDatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}

	// Setup a context channel to handle OS interrupts
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// Initialize the application with the database connection
	app := application.New(appDatabase)

	// Start the app
	app.Start(ctx)
}
