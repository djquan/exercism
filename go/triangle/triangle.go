// Package triangle provides a method to determine the kind of triangle
package triangle

import "math"

// Kind represents what a type of triangle
type Kind int

const (
	// NaT is not a triangle
	NaT = iota
	// Equ is Equilateral
	Equ
	// Iso is isosceles
	Iso
	// Sca is scalene
	Sca
)

// KindFromSides returns a Kind of triangle
func KindFromSides(a, b, c float64) Kind {
	for _, x := range []float64{a, b, c} {
		if x <= 0 || math.IsNaN(x) || math.IsInf(x, 0) {
			return NaT
		}
	}

	if a+b < c || a+c < b || b+c < a {
		return NaT
	}

	if a == b && b == c {
		return Equ
	}

	if a == b || a == c || b == c {
		return Iso
	}

	return Sca
}
