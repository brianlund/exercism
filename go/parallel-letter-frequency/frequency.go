package letter

// FreqMap maps letters and their frequency
type FreqMap map[rune]int

// Res holds the results from concurrent goroutines
type Res struct {
	Str    string
	ResMap FreqMap
}

// Frequency counts the frequency of letters in the input string
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency executes Frequency with goroutines
// result from calling Frequency is returned through a channel and a loop populates Res which is then returned
func ConcurrentFrequency(s []string) FreqMap {

	c := make(chan FreqMap)
	res := FreqMap{}
	for _, i := range s {
		go func(i string) {
			c <- Frequency(i)
		}(i)
	}
	for range s {
		for letter, count := range <-c {
			res[letter] += count
		}
	}
	return res
}
