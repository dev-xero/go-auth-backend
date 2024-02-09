package util

import (
	"log"
	"regexp"
)

/*
Compares and verifies if a provided string is a valid email format

Objectives:
  - Create a valid email regex pattern
  - Compare the provided string to the email regex pattern

Params:
  - email: A string to compare against the pattern

Returns:
  - True if the email string is a valid email pattern, false otherwise
*/
func IsValidEmail(email string) bool {
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compare the email string with the regex pattern
	regex, err := regexp.Compile(emailPattern)
	if err != nil {
		log.Println("[FAIL]: error compiling regex:", err)
		return false
	}

	return regex.MatchString(email)
}
