package luhn

import (
	"strconv"
	"strings"
)

//Valid given a number determine whether or not it is valid per the Luhn formula
func Valid(text string) bool {

	stp := strings.TrimSpace(text)

	s := strings.Replace(stp, " ", "", -1)

	if len(s) <= 1 {
		return false
	}

	var numbers []int

	for i := 0; i < len(s); i++ {
		num, err := strconv.ParseInt(string(s[i]), 10, 64)
		if err != nil {
			return false
		}
		numbers = append(numbers, int(num))
	}

	return doubleSum(numbers)%10 == 0
}

func doubleSum(num []int) int {

	var count int
	for i := 0; i < len(num); i++ {
		j := len(num) - i

		if j%2 == 0 {

			d := num[i] * 2

			if d > 9 {
				d -= 9
			}
			num[i] = d
		}

		count += num[i]
	}

	return count
}

//
