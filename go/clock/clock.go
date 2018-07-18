// Package clock provides functions for dealing with minutes on a clock
package clock

import (
	"fmt"
)

// Clock type maintains a representation of a clock
type Clock struct {
	h int
	m int
}

// New creates a new clock given hour and minutes
func New(h int, m int) Clock {
	h, m = normalize(h, m)
	return Clock{h, m}
}

// Add adds minutes onto a clock
func (c Clock) Add(m int) Clock {
	hour, minute := normalize(c.h, c.m+m)
	return Clock{hour, minute}
}

// Subtract removes minutes from a clock
func (c Clock) Subtract(m int) Clock {
	hour, minute := normalize(c.h, c.m-m)
	return Clock{hour, minute}
}

// String provides a string version of a clock
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

func normalize(h int, m int) (int, int) {
	for m >= 60 {
		m -= 60
		h++
	}

	for m < 0 {
		m += 60
		h--
	}

	for h >= 24 {
		h -= 24
	}

	for h < 0 {
		h += 24
	}

	return h, m
}
