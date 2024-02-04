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

func initDatabaseConnection() (*sql.DB, error) {
	// Attempt connecting to the database
	db, err := database.ConnectDatabase()
	if err != nil {
		msg := "[FAIL]: unable to connect database"
		return nil, fmt.Errorf("%s: %w", msg, err)
	}

	// Defer closing database connection if open
	defer func() {
		if db != nil {
			db.Close()
		}
	}()

	return db, nil
}

func main() {
	database, err := initDatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	app := application.New(database)
	app.Start(ctx)
}
