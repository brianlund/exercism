// Package raindrops solves the raindrop exercism
package raindrops

import (
	"strconv"
)

// Convert translates a number to raindrop-speak based on it's factores.
// If the numbers factors doesn't include any of the predefined factors
// it is returned as a string
func Convert(num int) string {

	raindrops := ""

	if num%3 == 0 {
		raindrops += "Pling"
	}
	if num%5 == 0 {
		raindrops += "Plang"
	}
	if num%7 == 0 {
		raindrops += "Plong"
	}

	if raindrops == "" {
		return strconv.Itoa(num)
	}

	return raindrops

}
