package validators

import (
	"log"

	"github.com/dev-xero/authentication-backend/util"
)

func ValidateUserInput(body *util.AuthRequestBody) error {
	if body.Username == "" || body.Email == "" || body.Password == "" {
		log.Println("[FAIL]: username, email or password is empty")
		return util.ErrEmptyFields
	}

	if !util.IsValidEmail(body.Email) {
		log.Println("[FAIL]: invalid email provided")
		return util.ErrEmailInvalid
	}

	// Password field must be at least 8 chars long
	if len(body.Password) < 8 {
		log.Println("[FAIL]: password field must contain at least 8 characters")
		return util.ErrPasswordLength
	}

	return nil
}
