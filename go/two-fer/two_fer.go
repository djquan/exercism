// Package twofer provides a simple function for two for one.
package twofer

import "fmt"

// ShareWith returns a sentence based on a passed in name
func ShareWith(n string) string {
	s := "you"
	if n != "" {
		s = n
	}

	return fmt.Sprintf("One for %s, one for me.", s)
}
