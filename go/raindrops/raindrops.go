// Package raindrops provides a method for converting a number into raindrop sounds
package raindrops

import (
	"sort"
	"strconv"
	"strings"
)

var dropMap = map[int]string{
	3: "Pling",
	5: "Plang",
	7: "Plong",
}

var keys = getKeys(dropMap)

// Convert converts a number into raindrop sounds based on factors
func Convert(i int) string {
	var s strings.Builder
	for _, k := range keys {
		if i%k == 0 {
			s.WriteString(dropMap[k])
		}
	}

	if s.Len() == 0 {
		s.WriteString(strconv.Itoa(i))
	}
	return s.String()
}

func getKeys(m map[int]string) []int {
	var keys []int
	for k := range dropMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}
