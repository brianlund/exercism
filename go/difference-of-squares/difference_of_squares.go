// Package diffsquares solves the Difference of Squares exercism problem
package diffsquares

// Difference calculares square of the sum and the sum of the squares of the first N natural numbers
func Difference(num int) int {
	return SquareOfSum(num) - SumOfSquares(num)
}

// SumOfSquares finds the sum of squares of the first N natural numbers
func SumOfSquares(num int) int {
	return (num + 1) * (num*2 + 1) * num / 6
}

// SquareOfSum finds the square of the sum of the first N natural numbers
func SquareOfSum(num int) (square int) {
	sum := (num + 1) * num / 2
	return sum * sum
}
