package route

import (
	"github.com/dev-xero/authentication-backend/handler"
	"github.com/go-chi/chi/v5"
)

func LoadAuthRoutes(router chi.Router) {
	auth := &handler.Auth{}

	router.Get("/", auth.Home)
	router.Post("/sign-up", auth.SignUp)
	router.Post("/sign-in", auth.SignIn)
	router.Post("/sign-out", auth.SignOut)
}