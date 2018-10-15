package romannumerals

import (
	"fmt"
	"strings"
)

var araToRom = []struct {
	arabic int
	roman  string
}{
	{1, "I"}, {4, "IV"}, {5, "V"}, {9, "IX"},
	{10, "X"}, {40, "XL"}, {50, "L"}, {90, "XC"},
	{100, "C"}, {400, "CD"}, {500, "D"}, {900, "CM"},
	{1000, "M"},
}

func ToRomanNumeral(arabic int) (string, error) {

	if arabic <= 0 || arabic > 3000 {
		return "", fmt.Errorf("number %d out of rangee, must be 0-3000", arabic)
	}
	var roman string
	for i := len(araToRom) - 1; i >= 0; i-- {

		if arabic-araToRom[i].arabic < 0 {
			continue
		}
		num := arabic / araToRom[i].arabic

		s := strings.Repeat(araToRom[i].roman, num)
		roman += s
		arabic -= num * araToRom[i].arabic
	}

	return roman, nil
}
