package util

import "strings"

/*
Capitalizes the first letter of a string

Objectives:
  - Return the string if its a single letter
  - Return the same string with the first letter capitalized

Params:
  - s: The string to capitalize

Returns:
  - The capitalized string
*/
func CapitalizeFirstLetter(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
