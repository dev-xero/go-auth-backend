package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetConnectionString() string {
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
