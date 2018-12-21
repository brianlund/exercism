// Package grains solves the exercism wheat grains on a chessboard problem
package grains

import "fmt"

// Square calculates the number of grains for a given chessboard square
func Square(square int) (grains uint64, err error) {

	if square <= 0 || square > 64 {
		return 0, fmt.Errorf("Square %v is outside of chessboard 8x8 size", square)
	}

	return 1 << (uint(square) - 1), nil

}

// Total calculates the total number of grains on a chessboard
func Total() uint64 {

	return 1<<64 - 1
}
