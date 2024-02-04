package util

import (
	"log"
	"regexp"
)

func IsValidEmail(email string) bool {
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	regex, err := regexp.Compile(emailPattern)
	if err != nil {
		log.Println("[FAIL]: error compiling regex:", err)
		return false
	}

	return regex.MatchString(email)
}
