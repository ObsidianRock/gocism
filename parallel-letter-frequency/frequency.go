package letter

type FreqMap map[rune]int

//Frequency counts the frequency of letters in texts using sequential computation
func Frequency(s string) FreqMap {

	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func breakdown(str []string, rc chan rune) {

	for i := 0; i < len(str); i++ {
		for _, j := range str[i] {
			rc <- j
		}
	}

	close(rc)
	return
}

//ConcurrentFrequency counts the frequency of letters in texts using parallel computation.
func ConcurrentFrequency(str []string) FreqMap {
	m := FreqMap{}

	rc := make(chan rune)

	go breakdown(str, rc)

	for i := range rc {
		m[i]++
	}

	return m
}
