// Package twofer provides functions for returning two-for-one
package twofer

import "fmt"

// ShareWith takes a name as input and returns two for one, for that name
func ShareWith(name string) string {
	var x string
	switch name {
	case "":
		x = "you"
	default:
		x = name
	}
	twofer := fmt.Sprintf("One for %s, one for me.", x)
	return twofer
}
