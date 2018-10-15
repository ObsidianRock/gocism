package atbash

import (
	"strings"
	"unicode"
)

var plain = "abcdefghijklmnopqrstuvwxyz"
var cipher = "zyxwvutsrqponmlkjihgfedcba"

func Atbash(text string) string {

	var norm string
	for _, v := range text {
		if unicode.IsSpace(v) || unicode.IsPunct(v) {
			continue
		}
		norm += strings.ToLower(string(v))
	}

	var ciph string
	for k, s := range norm {

		if k%5 == 0 && k != 0 {
			ciph += " "
		}

		i := strings.Index(plain, string(s))
		if i == -1 {
			ciph += string(s)
		} else {
			ciph += string(cipher[i])
		}

	}

	return ciph
}
