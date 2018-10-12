package cryptosquare

import (
	"math"
)

func processText(message string) string {
	var s []rune

	for _, v := range message {

		switch {
		case v >= 65 && v <= 90:
			v += 32
			s = append(s, v)
		case v >= 97 && v <= 122 || v >= 48 && v <= 57:
			s = append(s, v)
		default:
			continue
		}
	}

	//fmt.Println(string(s))
	return string(s)
}

func Encode(message string) string {

	newMessage := processText(message)
	c, r := square(len(newMessage))

	x := org(newMessage, c, r)
	s := code(x, c, r)

	return s
}

func square(length int) (int, int) {

	sqrt := math.Sqrt(float64(length))

	x := math.Remainder(float64(length), sqrt)

	var c, r int

	if x == 0.0 {
		c, r = int(sqrt), int(sqrt)

	} else {
		c = int(math.Round(sqrt + 0.5))

		r = int(math.Round(float64(length) / sqrt))
	}

	return c, r
}

func org(message string, c, r int) []string {

	var s []string

	total := c * r

	var i int

	//fmt.Println(c, r, c*r, len(message))
	for i < total {

		var mes string

		var j int

		for x := i; j < c; x++ {

			var m string
			if x >= len(message) {
				m = " "
			} else {
				m = string(message[x])
			}

			mes = mes + m
			i++
			j++
		}

		s = append(s, mes)
	}

	//fmt.Println(s)

	return s
}

func code(message []string, c, r int) string {

	var str string

	var i, chr int

	for chr < c*r {

		var xtr string

		for _, v := range message {

			xtr += string(v[i])
			chr++
		}
		str += xtr

		if chr < c*r {
			str += " "
		}

		i++
	}

	return str
}
