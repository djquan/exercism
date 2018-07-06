// Package raindrops provides a method for converting a number into raindrop sounds
package raindrops

import (
	"strconv"
	"strings"
)

// Convert converts a number into raindrop sounds based on factors
func Convert(i int) string {
	var s strings.Builder

	if i%3 == 0 {
		s.WriteString("Pling")
	}

	if i%5 == 0 {
		s.WriteString("Plang")
	}

	if i%7 == 0 {
		s.WriteString("Plong")
	}

	if s.Len() == 0 {
		s.WriteString(strconv.FormatInt(int64(i), 10))
	}
	return s.String()
}
