package collatzconjecture

import (
	"fmt"
)

// CollatzConjecture calculates the 3x+1 problem of any given positive integer
func CollatzConjecture(n int) (int, error) {
	steps := 0

	if n <= 0 {
		return 0, fmt.Errorf("Cannot calculate non positive numbers")
	}

	for {
		if n == 1 {
			break
		}

		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}

		steps++
	}
	return steps, nil
}
