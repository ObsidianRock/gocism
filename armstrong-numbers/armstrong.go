package armstrong

import (
	"math"
	"strconv"
)

func IsNumber(number int) bool {

	var sum int

	num := strconv.Itoa(number)
	factor := len(num)

	for _, v := range num {
		val, err := strconv.ParseInt(string(v), 10, 64)
		if err != nil {
			return false
		}

		sum += int(math.Pow(float64(val), float64(factor)))
	}

	return sum == number
}
