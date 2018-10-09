package etl

import (
	"strings"
)

//Transform ...
func Transform(input map[int][]string) map[string]int {

	var output = make(map[string]int)

	for k, v := range input {
		for _, s := range v {
			lower := strings.ToLower(s)
			output[lower] = k
		}
	}
	return output
}
