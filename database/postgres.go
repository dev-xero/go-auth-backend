package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

/*
Constructs a PostGreSQL connection string

Objectives:
  - Load private connection details as environment variables
  - Construct a connection string based on the details

Params:
  - No parameters

Returns:
  - A PostGreSQL connection string
*/
func getConnectionString() string {
	// Load environment variables from .env file in development
	if env := os.Getenv("ENVIRONMENT"); env != "production" {
		log.Println("[ENV]:", env)

		err := godotenv.Load()
		if err != nil {
			log.Fatal("[FATAL]: failed to connect database", err)
		}
	} else {
		log.Println("[ENV]: production")
	}

	// Store private connection details
	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_NAME")
	)

	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		user,
		password,
		host,
		port,
		dbname,
	)
}

/*
Connects to a PostGresQL database

Objectives:
  - Open a new database connection using the connection string
  - Ping the database to make sure the connection is open

Params:
  - No parameters

Returns:
  - A pointer to the SQL database
  - An error if unsuccessful
*/
func ConnectDatabase() (*sql.DB, error) {
	// Open a new database connection
	database, err := sql.Open("postgres", getConnectionString())

	// Return an error if the connection failed
	if err != nil {
		msg := "invalid connection string provided"
		log.Println("[FAIL]:", msg)

		return nil, fmt.Errorf(msg)
	}

	// Ping the database to verify the connection
	err = database.Ping()
	if err != nil {
		msg := "failed to establish a connection to the database"
		log.Println("[FAIL]:", msg)

		return nil, fmt.Errorf(msg)
	}

	log.Println("[SUCCESS]: connected database")

	return database, nil
}
