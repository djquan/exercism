package etl

import "strings"

//Transform takes a given map of points to letters and transforms it to a map of letters to points
func Transform(given map[int][]string) map[string]int {
	result := make(map[string]int)

	for point, letters := range given {
		for _, letter := range letters {
			result[strings.ToLower(letter)] = point
		}
	}

	return result
}
