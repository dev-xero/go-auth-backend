package service

import repository "github.com/dev-xero/authentication-backend/repository/user"

/*
AuthService handler struct

Objectives:
  - Handle all auth requests

Fields:
  - repo: The database repository
*/
type AuthService struct {
	Repo *repository.PostGreSQL
}

/*
Initializes a new auth service

Objectives:
  - Initialize an auth service with the provided repo

Params:
  - repo: The database repo to bind the handler to

Returns:
  - No return value
*/
func (auth *AuthService) New(repo *repository.PostGreSQL) {
	auth.Repo = repo
}
