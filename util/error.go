package util

import (
	"errors"
	"fmt"
)

type ValidationError error

// Stores all possible errors that may occur during input validation/
var (
	ErrEmptyFields    ValidationError = errors.New("one or more fields are empty")
	ErrEmailInvalid   ValidationError = errors.New("invalid email provided")
	ErrPasswordLength ValidationError = errors.New("password length must be at least 8 characters long")
)

/*
Returns a custom error format

Objectives:
  - Prefix the error with an error message

Params:
  - err: The error to format
  - msg: The string to prefix the error with

Returns:
  - A 64-bit signed integer, 0
  - A properly formatted error
*/
func Fail(err error, msg string) (int64, error) {
	return 0, fmt.Errorf("%s: %v", msg, err)
}
