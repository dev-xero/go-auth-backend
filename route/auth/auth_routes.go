package route

import (
	"database/sql"

	"github.com/dev-xero/authentication-backend/handler"
	repository "github.com/dev-xero/authentication-backend/repository/user"
	"github.com/go-chi/chi/v5"
)

func LoadAuthRoutes(router chi.Router, db *sql.DB) {
	auth := &handler.Auth{}
	auth.New(&repository.PostGreSQL{Database: db})

	router.Get("/", auth.Home)
	router.Post("/sign-up", auth.SignUp)
	router.Post("/sign-in", auth.SignIn)
	router.Post("/sign-out", auth.SignOut)
}
