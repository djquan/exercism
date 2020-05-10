package accumulate

//Accumulate maps over a given input with a given function
func Accumulate(given []string, converter func(string) string) []string {
	result := make([]string, 0, len(given))

	for _, val := range given {
		result = append(result, converter(val))
	}

	return result
}
