package application

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type App struct{}

func New() *App {
	app := &App{}
	return app
}

func baseHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API Home"))
}

func (app *App) Start(ctx context.Context) error {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("[FAIL]: unable to load environment variables")
	}

	var port string = os.Getenv("PORT")

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: http.HandlerFunc(baseHandler),
	}

	err = server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("[FAIL]: unable to start server")
	}

	return nil
}
