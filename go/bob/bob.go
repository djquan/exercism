// Package bob provides a Hey method for responding as a lackadaisical teeneager
package bob

import (
	"regexp"
	"strings"
)

// Hey returns a teenager's response based on the original response
func Hey(remark string) string {
	trimmedRemark := strings.Trim(remark, " \t\n\r")
	isQuestion := strings.HasSuffix(trimmedRemark, "?")
	hasLetters, _ := regexp.MatchString("[A-z]", trimmedRemark)

	switch {
	case trimmedRemark == "":
		return "Fine. Be that way!"
	case hasLetters && strings.ToUpper(trimmedRemark) == trimmedRemark:
		if isQuestion {
			return "Calm down, I know what I'm doing!"
		}
		return "Whoa, chill out!"
	case isQuestion:
		return "Sure."
	default:
		return "Whatever."
	}
}
