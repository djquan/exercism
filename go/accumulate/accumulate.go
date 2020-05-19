package accumulate

//Accumulate maps over a given input with a given function
func Accumulate(given []string, converter func(string) string) []string {
	result := make([]string, len(given))

	for i, val := range given {
		result[i] = converter(val)
	}

	return result
}
