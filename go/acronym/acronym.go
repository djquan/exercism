// Package acronym provdes a function that gives an acronym of word/words
package acronym

import (
	"strings"
)

// Abbreviate returns the acronym of a given string
func Abbreviate(s string) string {
	var strBuilder strings.Builder
	for _, w := range strings.FieldsFunc(s, toSplitAt) {
		runes := []rune(w)
		strBuilder.WriteRune(runes[0])
	}
	return strings.ToUpper(strBuilder.String())
}

func toSplitAt(r rune) bool {
	return r == ' ' || r == '-'
}
