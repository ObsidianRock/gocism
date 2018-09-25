package isogram

import (
	"strings"
)

//IsIsogram determines if a word or phrase is an isogram.
func IsIsogram(text string) bool {

	if len(text) == 0 {
		return true
	}

	lowerW := strings.ToLower(text)

	for i, r := range lowerW {
		if string(r) == " " || string(r) == "-" {
			continue
		}
		if i < strings.LastIndex(lowerW, string(r)) {
			return false
		}
	}
	return true

}
