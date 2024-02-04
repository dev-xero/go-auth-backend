package application

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dev-xero/authentication-backend/route"
	"github.com/joho/godotenv"
)

type App struct {
	router   http.Handler
	database *sql.DB
}

func New(db *sql.DB) *App {
	app := &App{
		router:   route.LoadRoutes(db),
		database: db,
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

	// Handle server listening on port in a goroutine
	go func() {
		err = server.ListenAndServe()
		if err != nil {
			msg := "[FAIL]: unable to start server"

			errorChan <- fmt.Errorf("%s: %w", msg, err)
			close(errorChan)
		}
	}()

	url := fmt.Sprintf("http://127.0.0.1:%s", port)
	log.Printf("[SUCCESS]: server listening at: %s\n", url)

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
