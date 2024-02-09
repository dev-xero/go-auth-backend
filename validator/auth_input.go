package validators

import (
	"log"

	"github.com/dev-xero/authentication-backend/util"
)

/*
Validates user input from the auth request body.

Objectives:
  - Request body must contain non-empty username, email and password
  - The provided email must be valid
  - Password length must be at least 8 characters

Params:
  - Auth request body with

Returns:
  - Returns an error if the auth request body doesn't pass all checks
*/
func ValidateUserInput(body *util.AuthRequestBody) error {
	if body.Username == "" || body.Email == "" || body.Password == "" {
		log.Println("[FAIL]: username, email or password is empty")
		return util.ErrEmptyFields
	}

	if !util.IsValidEmail(body.Email) {
		log.Println("[FAIL]: invalid email provided")
		return util.ErrEmailInvalid
	}

	if len(body.Password) < 8 {
		log.Println("[FAIL]: password field must contain at least 8 characters")
		return util.ErrPasswordLength
	}

	return nil
}
