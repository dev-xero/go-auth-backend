package util

import (
	"log"
	"golang.org/x/crypto/bcrypt"
)

/*
* Options for generating the hash
* params:
	* - Min: minimum allowable cost
	* - Max: maximum allowable cost
	* - Base: default cost to use if the the cost passed in is below Min
*/
type HashCost struct {
	Min  int
	Max  int
	Base int
}

/*
* Generates a bcrypt hash of a string argument with the provided options
 */
func GenerateHash(str string, costOptions HashCost) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(str), costOptions.Base)
	if err != nil {
		log.Println("[FAIL]: could not generate hash")
		return "", err
	}

	return string(hashedPassword), nil
}
