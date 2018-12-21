// Package isogram provides functions for determining if a word is an Isogram
package isogram

import (
	"unicode"
)

// IsIsogram returns true if the input has no repeating characters excluding "-" and blank space.
func IsIsogram(isoword string) bool {
	seen := map[rune]bool{}

	for _, r := range isoword {
		if !unicode.IsLetter(r) {
			continue
		}
		lower := unicode.ToLower(r)
		if _, ok := seen[lower]; ok {
			return false
		}
		seen[lower] = true
	}
	return true
}
