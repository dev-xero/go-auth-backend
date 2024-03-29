package route

import (
	"database/sql"

	handler "github.com/dev-xero/authentication-backend/handler/auth"
	repository "github.com/dev-xero/authentication-backend/repository/user"
	"github.com/dev-xero/authentication-backend/service"
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
	authDBService := &service.DatabaseProvider{}
	authDBService.New(&repository.PostGreSQL{Database: db})

	authHandler := &handler.AuthHandler{}
	authHandler.WithService(authDBService)

	router.Get("/", authHandler.Home)
	router.Post("/sign-up", authHandler.SignUp)
	router.Post("/sign-in", authHandler.SignIn)
	router.Post("/sign-out", authHandler.SignOut)
	router.Get("/oauth/google", authHandler.GoogleSignIn)
	router.Get("/oauth/google/callback", authHandler.GoogleSignInCallback)
	router.Get("/oauth/{x}/failure", authHandler.OAuthFailure)
}
