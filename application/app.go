package application

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dev-xero/authentication-backend/database"
	"github.com/dev-xero/authentication-backend/route"
	"github.com/joho/godotenv"
)

type App struct {
	router http.Handler
}

func New() *App {
	app := &App{
		router: route.LoadRoutes(),
	}

	return app
}

func (app *App) Start(ctx context.Context) error {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("[FAIL]: unable to load environment variables")
	}

	var port string = os.Getenv("PORT")

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: app.router,
	}

	errorChan := make(chan error, 1)

	// Attempt connecting to the database
	db, err := database.ConnectDatabase()
	if err != nil {
		msg := "[FAIL]: unable to connect database"

		errorChan <- fmt.Errorf("%s", msg)
		close(errorChan)
	}

	// Close database connection if open
	defer func() {
		if db != nil {
			db.Close()
		}
	}()

	// Handle server listening on port in a goroutine
	go func() {
		err = server.ListenAndServe()
		if err != nil {
			msg := "[FAIL]: unable to start server"

			errorChan <- fmt.Errorf("%s: %w", msg, err)
			close(errorChan)
		}
	}()

	// Handle graceful termination
	select {
	case err = <-errorChan:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		return server.Shutdown(timeout)
	}
}
