package letter

import (
	"sync"
)

var wg sync.WaitGroup

func ConcurrentFrequency(letters []string) FreqMap {
	wg.Add(len(letters))
	m := make([]FreqMap, 0, len(letters))

	for _, letter := range letters {
		go func(l string) {
			m = append(m, Frequency(l))
			wg.Done()
		}(letter)
	}

	wg.Wait()

	mm := FreqMap{}

	for i := range m {
		for k, v := range m[i] {
			mm[rune(k)] += v
		}
	}

	return mm
}
