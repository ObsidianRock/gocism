package acronym

import (
	"strings"
)

//Abbreviate converts a phrase to its acronym.
func Abbreviate(s string) string {

	rs := strings.Replace(s, "-", " ", -1)
	sp := strings.Split(rs, " ")

	var acr string

	for _, v := range sp {
		acrU := strings.ToUpper(string(v[0]))
		acr = acr + acrU
	}

	return acr
}
