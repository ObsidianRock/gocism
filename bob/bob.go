package bob

import (
	"strings"
)

//Hey ...
func Hey(remark string) string {

	trimmed := removeWhite(remark)

	if len(trimmed) == 0 {
		return "Fine. Be that way!"
	}

	lastCharacter := lastChar(trimmed)
	allUpperCase := allBlock(trimmed)

	if lastCharacter == "?" {

		if allUpperCase {
			return "Calm down, I know what I'm doing!"
		}

		return "Sure."
	}

	if lastCharacter == "?" && allUpperCase {
		return "Calm down, I know what I'm doing!"
	}

	if lastCharacter == "!" || allUpperCase {

		return "Whoa, chill out!"
	}

	return "Whatever."

}

func allBlock(s string) bool {
	return strings.ToUpper(s) == s && strings.ToLower(s) != strings.ToUpper(s)
}

func lastChar(s string) string {
	return string(s[len(s)-1])
}

func removeWhite(s string) string {
	trimmed := strings.TrimSpace(s)
	return strings.Replace(trimmed, " ", "", -1)
}
