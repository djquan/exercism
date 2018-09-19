// Package scrabble provides a function for calculating the scrabble score
package scrabble

import "unicode"

var letterScore = map[rune]int{
	'A': 1, 'B': 3, 'C': 3,
	'D': 2, 'E': 1, 'F': 4,
	'G': 2, 'H': 4, 'I': 1,
	'J': 8, 'K': 5, 'L': 1,
	'M': 3, 'N': 1, 'O': 1,
	'P': 3, 'Q': 10, 'R': 1,
	'S': 1, 'T': 1, 'U': 1,
	'V': 4, 'W': 4, 'X': 8,
	'Y': 4, 'Z': 10,
}

// Score calculates a scrabble score of a given word
func Score(w string) int {
	s := 0
	for _, l := range w {
		s += letterScore[unicode.ToUpper(l)]
	}
	return s
}
