package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func getConnectionString() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("[FATAL]: failed to connect database", err)
	}

	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_NAME")
	)

	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname,
	)
}

func ConnectDatabase() (*sql.DB, error) {
	database, err := sql.Open("postgres", getConnectionString())

	if err != nil {
		msg := "invalid connection string provided"
		log.Println("[FAIL]:", msg)

		return nil, fmt.Errorf(msg)
	}

	err = database.Ping()
	if err != nil {
		msg := "failed to establish a connection to the database"
		log.Println("[FAIL]:", msg)

		return nil, fmt.Errorf(msg)
	}

	log.Println("[SUCCESS]: connected database")

	return database, nil
}
