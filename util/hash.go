package util

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

/*
Options for generating the hash

Params:
  - Min: minimum allowable cost
  - Max: maximum allowable cost
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

Params:
  - str: the string to hash
  - costOptions: contains configurations for the hashing cost

Returns:
  - a string which is the hash
  - an error if the hashing failed
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

Params:
  - hash: the hash to compare against
  - str: the string to compare with the hash

Returns:
  - true if the string matches the hash, false otherwise
*/
func CompareWithHash(hash []byte, str string) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(str))
	return err == nil
}
