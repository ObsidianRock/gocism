package grains

import (
	"fmt"
	"math"
)

//Square calculates the number of grains of wheat on a chessboard
func Square(i int) (uint64, error) {

	if i <= 0 || i > 64 {
		return 0, fmt.Errorf("not valid number: %d,", i)
	}

	return uint64(math.Pow(float64(2), float64(i-1))), nil

}

//Total the total number of grains on a chessboard
func Total() uint64 {

	var total uint64

	for i := 0; i < 64; i++ {
		total += uint64(math.Pow(float64(2), float64(i)))
	}

	return total
}
