// Package scrabble computes the scrabble scores
package scrabble

import "unicode"

// Score returns the scrabble score for a given word
func Score(scrabbleword string) int {

	score := 0
	for _, c := range scrabbleword {
		switch unicode.ToLower(c) {
		case 'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't':
			score++
		case 'd', 'g':
			score += 2
		case 'b', 'c', 'm', 'p':
			score += 3
		case 'f', 'h', 'v', 'w', 'y':
			score += 4
		case 'k':
			score += 5
		case 'j', 'x':
			score += 8
		case 'q', 'z':
			score += 10
		}
	}
	return score

}
