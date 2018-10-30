package wordy

import (
	"math"
)


const ( 
	add = iota + 1
	sub 
	multi
	devide 
	power
)

type arithmetic struct {
	n1, n2 int
	op     string

}

//Answer Parse and evaluate simple math word problems returning the answer as an integer.
func Answer(q string) (int, bool) {
}


func op(n1, n2 int, op string) (n3, bool) {
	switch op{
	case 1:
		return n1 + n2 , true 
	case 2:
		return n1 - n2, true 
	case 3:
		return n1 * n2, true 
	case 4:
		return n1 / n2, true 
	case 5:
		return int(math.Pow(float64(n1), float64(n2)))
	}
}	

func parse(q string) []arithmetic, error {

}
