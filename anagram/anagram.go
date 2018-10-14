package anagram

import (
	"strings"
)

func Detect(subject string, candidates []string) []string {

	var anagrams []string
	for _, v := range candidates {

		anagram := isAnagram(strings.ToLower(subject), strings.ToLower(v))

		if !anagram {
			continue
		}

		anagrams = append(anagrams, v)
	}

	return anagrams
}

func isAnagram(subject, candidate string) bool {

	if len(candidate) != len(subject) {
		return false
	}

	if subject == candidate {
		return false
	}

	var sub = make(map[rune]int)
	var can = make(map[rune]int)

	for _, v := range subject {
		sub[v]++
	}

	for _, v := range candidate {
		can[v]++
	}

	for _, v := range candidate {

		count, ok := sub[v]

		if !ok {
			return false
		}

		if count != can[v] {
			return false
		}
	}

	return true
}
