package util

import "github.com/mrz1836/go-sanitize"

type AuthRequestBody struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SanitizeUserInput(body *AuthRequestBody) {
	body.Email = sanitize.Email(body.Email, false)
	body.Username = sanitize.Alpha(body.Username, false)
	body.Password = sanitize.AlphaNumeric(body.Password, false)
}
