// Package luhn provides a method for determining if a number is a valid credit card number
package luhn

import (
	"strings"
)

// Valid checks a given string and returns true if it is luhn valid
func Valid(n string) bool {
	n = strings.Replace(n, " ", "", -1)
	if len(n) <= 1 {
		return false
	}

	double := false
	sum := 0
	for i := len(n) - 1; i >= 0; i-- {
		digit := int(n[i] - '0')
		if digit > 9 {
			return false
		}

		sum += digit

		if double {
			sum += digit
			if digit*2 > 9 {
				sum -= 9
			}
		}

		double = !double
	}

	return sum%10 == 0
}
