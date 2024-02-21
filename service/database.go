package service

import repository "github.com/dev-xero/authentication-backend/repository/user"

/*
DatabaseProvider handler struct

Objectives:
  - Provides database repositories to handlers

Fields:
  - repo: The database repository
*/
type DatabaseProvider struct {
	Repo *repository.PostGreSQL
}

/*
Initializes a new database service

Params:
  - repo: The database repo to bind the handler to

Returns:
  - No return value
*/
func (database *DatabaseProvider) New(repo *repository.PostGreSQL) {
	database.Repo = repo
}
