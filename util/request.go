package util

import "github.com/mrz1836/go-sanitize"

/*
Sanitizable Interface

Defines structs that implement the sanitize function
*/
type Sanitizable interface {
	Sanitize()
}

/*
Sign-up auth request body

Fields:
  - Username: string
  - Email:    string
  - Password: string
*/
type SignUpRequestBody struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Implement sanitize function for the sign-up request body
func (sanitizable *SignUpRequestBody) Sanitize() {
	sanitizable.Email = sanitize.Email(sanitizable.Email, false)
	sanitizable.Username = sanitize.Alpha(sanitizable.Username, false)
	sanitizable.Password = sanitize.AlphaNumeric(sanitizable.Password, false)
}

/*
Sign-in auth request body

Fields:
  - Email:    string
  - Password: string
*/
type SignInRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Implement the sanitize function fo the sign-in request body
func (sanitizable *SignInRequestBody) Sanitize() {
	sanitizable.Email = sanitize.Email(sanitizable.Email, false)
	sanitizable.Password = sanitize.AlphaNumeric(sanitizable.Password, false)
}

/*
Sanitize user input from request body

Objectives:
  - Sanitize the email
  - Sanitize the username
  - Sanitize the password

Params:
  - body: A sanitizable request body

Returns:
  - No return value
*/
func SanitizeUserInput(body Sanitizable) {
	body.Sanitize()
}
