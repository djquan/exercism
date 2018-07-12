/*
Package grains provides functions for calculating grain
based on grain doubling on each square.
*/
package grains

import "errors"

/*
Square calculates the number of grains on a square n if
one were to double the number of grains on each square.
*/
func Square(n int) (uint64, error) {
	if n <= 0 {
		return 0, errors.New("Cannot square a number less than 1")
	}

	if n > 64 {
		return 0, errors.New("Cannot square a number over 64")
	}

	g := uint64(1)

	for i := 2; i <= n; i++ {
		g = g * 2
	}

	return g, nil
}

/*
Total calculates the number of grains on a chessboard if
one were to double the number of grains on each square.
*/
func Total() uint64 {
	total := uint64(1)
	square := uint64(1)

	for i := 2; i <= 64; i++ {
		square = square * 2
		total += square
	}

	return total
}
