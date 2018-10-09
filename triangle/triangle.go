package triangle

import (
	"math"
)

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	// Pick values for the following identifiers used by the test program.
	NaT = iota
	Equ // equilateral
	Iso // isosceles
	Sca // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {

	switch {
	case !isTriangle(a, b, c):
		return NaT
	case a == b && a == c:
		return Equ
	case a == b || a == c || b == c:
		return Iso
	}

	return Sca
}

func isTriangle(a, b, c float64) bool {

	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}

	if a == math.Inf(1) || b == math.Inf(1) || c == math.Inf(1) {
		return false
	}

	var biggest, side1, side2 float64

	switch {
	case a > b && a > c:
		biggest = a
		side1 = b
		side2 = c
	case b > a && b > c:
		biggest = b
		side1 = a
		side2 = c
	default:
		biggest = c
		side1 = a
		side2 = b
	}

	if !(side1+side2 >= biggest) {
		return false
	}

	return true
}
