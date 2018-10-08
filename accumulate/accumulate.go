package accumulate

func Accumulate(words []string, converter func(string) string) []string {

	for i := 0; i < len(words); i++ {
		words[i] = converter(words[i])
	}

	return words
}
