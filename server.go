package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/dev-xero/authentication-backend/application"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	app := application.New()
	app.Start(ctx)
}
