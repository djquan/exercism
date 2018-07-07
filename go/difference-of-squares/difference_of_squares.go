// Package diffsquares provides functions for operating on the square of sums and sum of squares
package diffsquares

// SumOfSquares calculates the sum of squares up to a given number
func SumOfSquares(n int) int {
	total := 0
	for i := 0; i <= n; i++ {
		total += (i * i)
	}
	return total
}

// SquareOfSums calculates the square of sums up to a given number
func SquareOfSums(n int) int {
	total := 0
	for i := 0; i <= n; i++ {
		total += i
	}
	return total * total
}

// Difference calculates the difference between the square of sums and sum of squares
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
