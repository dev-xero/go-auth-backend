package application

import (
	"context"
	"fmt"
	"net/http"
	"os"

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

	err = server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("[FAIL]: unable to start server")
	}

	return nil
}
