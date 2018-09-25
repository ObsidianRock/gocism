package raindrops

import "fmt"

//Convert converts a number to a string based on numbers factor
func Convert(num int) string {

	var raindrops string

	if num%3 != 0 && num%5 != 0 && num%7 != 0 {
		return fmt.Sprintf("%d", num)
	}

	if num%3 == 0 {
		raindrops += "Pling"
	}

	if num%5 == 0 {
		raindrops += "Plang"
	}

	if num%7 == 0 {
		raindrops += "Plong"
	}

	return raindrops
}
