package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency calculates frquency of runes over multile strings concurrently
func ConcurrentFrequency(inputList []string) FreqMap {
	m := FreqMap{}
	c := make(chan FreqMap, 5)

	for _, input := range inputList {
		go func(inputString string) {
			c <- Frequency(inputString)
		}(input)
	}

	for range inputList {
		freqMap := <-c
		for i := range freqMap {
			m[i] += freqMap[i]
		}
	}
	return m
}
