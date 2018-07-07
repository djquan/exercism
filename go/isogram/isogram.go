// Package isogram provides a function for checking if a word is an isogram
package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram checks a given word and returns true if it is an isogram
func IsIsogram(s string) bool {
	letters := make(map[rune]bool)
	for _, r := range strings.ToUpper(s) {
		if !unicode.IsLetter(r) {
			continue
		}

		if letters[r] {
			return false
		}
		letters[r] = true
	}

	return true
}
