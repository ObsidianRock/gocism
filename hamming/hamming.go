package hamming

import (
	"fmt"
)

//Distance calculates the Hamming difference between two DNA strands.
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return 0, fmt.Errorf("DNA strands do not match a:%d, b:%d", len(a), len(b))
	}

	var hamming int

	for i := 0; i < len(a); i++ {

		if a[i] != b[i] {
			hamming++
		}
	}

	return hamming, nil
}
