package letter

type FreqMap map[rune]int

//Frequency ...
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

//ConcurrentFrequency ...
func ConcurrentFrequency(str []string) FreqMap {
	m := FreqMap{}

	rc := make(chan rune)

	go breakdown(str, rc)

	for i := range rc {
		m[i]++
	}

	return m
}
