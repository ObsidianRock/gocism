package lsproduct

import (
	"fmt"
	"strconv"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {

	if span > len(digits) || span < 0 {
		return -1, fmt.Errorf("length digits: %d, span: %d", len(digits), span)
	}

	var nums []int

	for _, v := range digits {
		i, err := strconv.Atoi(string(v))
		if err != nil {
			return -1, fmt.Errorf("Not a number")
		}

		nums = append(nums, i)
	}

	end, front := 0, span
	var biggest int64

	for front <= len(digits) {

		n := nums[end:front]

		x := 1
		for _, v := range n {
			x = x * v
		}

		if int64(x) > biggest {
			biggest = int64(x)
		}

		front++
		end++
	}

	return biggest, nil
}
