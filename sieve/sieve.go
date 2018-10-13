package sieve

import (
	"math"
)

func Sieve(lim int) []int {

	if lim <= 1 {
		return []int{}
	}

	var arr = make([]bool, lim+1)

	for i := 2; i < len(arr); i++ {
		arr[i] = true
	}

	sqrt := int(math.Sqrt(float64(lim)))

	for i := 2; i <= sqrt; i++ {
		if arr[i] == true {
			b := 0
			for j := (i * i); j <= lim; j = (i * i) + b*i {
				arr[j] = false
				b++
			}
		}
	}

	var prime []int
	for k, v := range arr {
		if v == true {
			prime = append(prime, k)
		}
	}

	return prime
}
