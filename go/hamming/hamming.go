// Package hamming provides a function for calculating hamming distance
package hamming

import "errors"

// Distance calculates hamming distances between two strings
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Lengths differ")
	}

	d := 0
	for i, r := range a {
		if rune(b[i]) != r {
			d++
		}
	}
	return d, nil
}
