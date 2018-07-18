package letter

// ConcurrentFrequency calculates the letter frequency in an array of strings
func ConcurrentFrequency(input []string) FreqMap {
	c := make(chan FreqMap)

	for _, t := range input {
		go func(t string) {
			c <- Frequency(t)
		}(t)
	}

	freqMap := FreqMap{}

	for range input {
		for k, v := range <-c {
			freqMap[k] += v
		}
	}

	return freqMap
}
