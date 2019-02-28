package letter

// This uses a map as a histogram to calculate the frequency of occurences
func ConcurrentFrequency(input []string) FreqMap {

	char := make(chan FreqMap)

	for _, s := range input {
		go func(s string) {
			char <- Frequency(s)
		}(s)
	}

	output := <-char

	for i := 1; i < len(input); i++ {
		for k, v := range <-char {
			output[k] += v
		}
	}

	return output
}