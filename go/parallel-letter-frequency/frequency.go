package letter

// FreqMap is a type that represents rune frequency in a string
type FreqMap map[rune]int

// Frequency calculates the letter frequency in a string
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}
