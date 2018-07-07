// Package reverse provides a function for reversing a string
package reverse

// String reverses a given String.
func String(s string) string {
	in := []rune(s)
	out := make([]rune, len(in))

	for i, r := range in {
		out[len(in)-i-1] = r
	}

	return string(out)
}
