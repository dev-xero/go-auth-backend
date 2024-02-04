package main

import (
	"context"

	"github.com/dev-xero/authentication-backend/application"
)

func main() {
	app := application.New()
	app.Start(context.Background())
}
