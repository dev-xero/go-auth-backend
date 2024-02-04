package validators

import (
	"fmt"

	"github.com/dev-xero/authentication-backend/util"
)

type body struct {
	Username string
	Email    string
	Password string
}

func ValidateUserInput(body *body) error {
	// Non of the fields must be empty
	if body.Username == "" || body.Email == "" || body.Password == "" {
		return fmt.Errorf("[FAIL]: username, email or password is empty")
	}

	// The email has to be valid
	if !util.IsValidEmail(body.Email) {
		return fmt.Errorf("[FAIL]: invalid email provided")
	}

	// Password field must be at least 8 chars long
	if len(body.Password) < 8 {
		return fmt.Errorf("[FAIL]: password field must contain at least 8 characters")
	}

	return nil
}
