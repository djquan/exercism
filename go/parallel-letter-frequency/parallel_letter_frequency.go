package letter

import (
	"sync"
)

// ConcurrentFrequency calculates the letter frequency in an array of strings
func ConcurrentFrequency(input []string) FreqMap {
	freqMap := FreqMap{}
	mutex := &sync.Mutex{}
	wg := &sync.WaitGroup{}

	for _, s := range input {
		wg.Add(1)
		go func(s string) {
			for _, r := range s {
				mutex.Lock()
				freqMap[r]++
				mutex.Unlock()
			}
			wg.Done()
		}(s)
	}

	wg.Wait()
	return freqMap
}
