// Package clock provides functions for dealing with minutes on a clock
package clock

import (
	"fmt"
)

// Clock type maintains a representation of a clock
type Clock int

const minutesInDay = 1440

// New creates a new clock given hour and minutes
func New(h int, m int) Clock {
	c := (h*60 + m) % minutesInDay

	for c < 0 {
		c += minutesInDay
	}

	return Clock(c)
}

// Add adds minutes onto a clock
func (c Clock) Add(m int) Clock {
	return New(0, int(c)+m)
}

// Subtract removes minutes from a clock
func (c Clock) Subtract(m int) Clock {
	return New(0, int(c)-m)
}

// String provides a string version of a clock
func (c Clock) String() string {
	h := int(c) / 60
	m := int(c) % 60
	return fmt.Sprintf("%02d:%02d", h, m)
}
