package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(input string) Frequency {

	freg := make(Frequency)

	var words []rune
	space := true
	for _, v := range input {

		if unicode.IsPunct(v) {

			if v == '\'' {
				words = append(words, v)
				space = false
				continue
			}

			words = append(words, rune(32))

			if space {
				space = false
				continue
			}

			space = true
			continue
		}

		if unicode.IsSpace(v) {
			if !space {
				words = append(words, rune(32))
				space = true
			}
			continue
		}

		if unicode.IsLetter(v) || unicode.IsNumber(v) {
			words = append(words, v)
			space = false
		}
	}

	split := strings.Split(string(words), " ")
	for _, val := range split {

		if len(val) == 0 {
			continue
		}

		if val[0] == '\'' || val[len(val)-1] == '\'' {
			val = strings.Replace(string(val), "'", "", 2)
		}

		freg[strings.ToLower(val)]++
	}

	return freg
}
