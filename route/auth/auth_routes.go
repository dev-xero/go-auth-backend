package route

import (
	"database/sql"

	"github.com/dev-xero/authentication-backend/handler"
	repository "github.com/dev-xero/authentication-backend/repository/user"
	"github.com/go-chi/chi/v5"
)

/*
Loads auth routes

Objectives:
  - Setup an auth sub-router
  - Setup a database repository
  - Handle requests made to auth routes

Params:
  - router: A chi router
  - db:     A pointer to the application database

Returns:
  - No return value
*/
func LoadAuthRoutes(router chi.Router, db *sql.DB) {
	auth := &handler.Auth{}
	auth.New(&repository.PostGreSQL{Database: db})

	router.Get("/", auth.Home)
	router.Post("/sign-up", auth.SignUp)
	router.Post("/sign-in", auth.SignIn)
	router.Post("/sign-out", auth.SignOut)
}
