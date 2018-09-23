// Package scrabble provides a function for calculating the scrabble score
package scrabble

import "unicode"

// Score calculates a scrabble score of a given word
func Score(w string) int {
	s := 0
	for _, l := range w {
		switch unicode.ToUpper(l) {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			s++
		case 'D', 'G':
			s += 2
		case 'B', 'C', 'M', 'P':
			s += 3
		case 'F', 'H', 'V', 'W', 'Y':
			s += 4
		case 'K':
			s += 5
		case 'J', 'X':
			s += 8
		case 'Q', 'Z':
			s += 10
		}
	}
	return s
}
