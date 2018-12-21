// Package hamming provices a function for calculating the Hamming distance between DNA strands
package hamming

import "fmt"

// Distance calculates the Hamming distance between string a and string b.Distance
// It returns the distance and any error encountered
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return 0, fmt.Errorf("Unequal length strings not supported, %q and %q", a, b)
	}

	hamming := 0

	for i := range a {
		if a[i] != b[i] {
			hamming++
		}
	}
	return hamming, nil

}
