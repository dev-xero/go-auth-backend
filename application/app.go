package application

import (
	"context"
	"fmt"
	"net/http"
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
	server := &http.Server{
		Addr:    ":3000",
		Handler: http.HandlerFunc(baseHandler),
	}

	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("[FAIL]: unable to start server")
	}

	return nil
}
