package route

import (
	"database/sql"

	"github.com/dev-xero/authentication-backend/handler"
	"github.com/dev-xero/authentication-backend/middleware"
	repository "github.com/dev-xero/authentication-backend/repository/user"
	"github.com/go-chi/chi/v5"
)

/*
Loads user routes

Objectives:
  - Setup a user sub-router
  - Setup a database repository
  - Handle requests made to user routes

Params:
  - router: A chi router
  - db:     A pointer to the application database

Returns:
  - No return value
*/
func LoadUserRoutes(router chi.Router, db *sql.DB) {
	user := &handler.User{}
	user.New(&repository.PostGreSQL{Database: db})

	router.Get("/", user.Home)
	router.With(middleware.AuthenticateMiddleware).Get("/{id}", user.GetUserByID)
}
