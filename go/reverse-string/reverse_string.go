// Package reverse provides a function for reversing a string
package reverse

// String reverses a given String.
func String(s string) string {
	out := []rune(s)
	n := len(out)

	for i := 0; i < n/2; i++ {
		j := n - 1 - i
		out[i], out[j] = out[j], out[i]
	}

	return string(out)
}
