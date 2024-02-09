package util

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

/*
Options for generating the hash

Params:
  - Min:  minimum allowable cost
  - Max:  maximum allowable cost
  - Base: default cost to use if the the cost passed in is below Min
*/
type HashCost struct {
	Min  int
	Max  int
	Base int
}

var DefaultHashCost = HashCost{
	Min:  10,
	Max:  14,
	Base: bcrypt.DefaultCost,
}

/*
Generates a bcrypt hash of a string argument with the provided options

Objectives:
  - Generate a bcrypt hash from the password using a base cost

Params:
  - str:         The string to hash
  - costOptions: Contains configurations for the hashing cost

Returns:
  - A string which is the hash
  - An error if the hashing failed
*/
func GenerateHash(str string, costOptions HashCost) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(str), costOptions.Base)
	if err != nil {
		log.Println("[FAIL]: could not generate hash")
		return "", err
	}

	return string(hashedPassword), nil
}

/*
Compares the hash with the string, returns true if they match

Objectives:
  - Compare the provided string with a hash
  - Evaluate whether the string is equivalent to the hash

Params:
  - hash: The hash to compare against
  - str:  The string to compare with the hash

Returns:
  - True if the string matches the hash, false otherwise
*/
func CompareWithHash(hash []byte, str string) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(str))
	return err == nil
}
