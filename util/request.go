package util

import "github.com/mrz1836/go-sanitize"

/*
Auth request body

Fields:
  - Username: string
  - Email:    string
  - Password: string
*/
type AuthRequestBody struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

/*
Sanitize user input from request body

Objectives:
  - Sanitize the email
  - Sanitize the username
  - Sanitize the password

Params:
  - body: Auth request body

Returns:
  - No return value
*/
func SanitizeUserInput(body *AuthRequestBody) {
	body.Email = sanitize.Email(body.Email, false)
	body.Username = sanitize.Alpha(body.Username, false)
	body.Password = sanitize.AlphaNumeric(body.Password, false)
}
