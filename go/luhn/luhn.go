// Package luhn solves the exercism luhn challenge
package luhn

import (
	"strings"
	"unicode"
)

// Valid is used to validate a string of numbers against the luhn formula
func Valid(luhn string) bool {

	// Strip spaces from string to be checked
	num := []rune(strings.Replace(luhn, " ", "", -1))

	// Strings of length 1 or less are not valid
	if len(num) <= 1 {
		return false
	}

	checksum := 0
	alternate := false
	n := 0

	for i := len(num) - 1; i >= 0; i-- {
		if !unicode.IsNumber(num[i]) {
			return false
		}
		n = int((num[i]) - '0')
		if alternate {
			n *= 2
			if n > 9 {
				n = n - 9
			}
		}
		checksum += n
		alternate = !alternate
	}
	return (checksum%10 == 0)
}
