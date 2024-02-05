package util

import (
	"errors"
	"fmt"
)

type ValidationError error

// Store possible validation errors that could arise
var (
	ErrEmptyFields    ValidationError = errors.New("one or more fields are empty")
	ErrEmailInvalid   ValidationError = errors.New("invalid email provided")
	ErrPasswordLength ValidationError = errors.New("password length must be at least 8 characters long")
)

func Fail(err error, msg string) (int64, error) {
	// Custom error logger
	return 0, fmt.Errorf("%s: %v", msg, err)
}
