package clock

import (
	"fmt"
	"math"
)

//Clock type
type Clock struct {
	sec int
}

//New ...
func New(hours, minutes int) Clock {

	var h, m int

	h, m = hours, minutes

	if hours < 0 {
		x := hours % 24
		// hours = 24 + (-h)
		h = 24 + x
	}

	if minutes < 0 {
		x := minutes % 1440
		// minutes in 24 hours
		m = 1440 + x
	}

	if hours > 24 {
		x := hours % 24
		h = x
	}

	if minutes > 1440 {
		x := minutes % 1440
		m = x
	}

	var secs int
	secs = (60 * 60 * h) + (60 * m)

	if secs > 86400 {
		s := secs % 86400
		secs = s
	}

	c := Clock{
		sec: secs,
	}

	return c
}

func (c Clock) String() string {

	x := int(math.Abs(float64(c.sec))) % 86400

	m := x / 60

	h := ((x / 60) / 60)

	min := m % 60

	hour := h % 24

	return fmt.Sprintf("%02d:%02d", hour, min)
}

//Add ...
func (c Clock) Add(minutes int) Clock {

	c.sec += (minutes * 60)

	return c
}

//Subtract ...
func (c Clock) Subtract(minutes int) Clock {

	if minutes*60 > 86400 {
		l := (minutes * 60) % 86400

		minutes = l / 60
	}

	if (c.sec - minutes*60) < 0 {

		m := 86400 + c.sec

		k := m - (minutes * 60)

		c.sec = k

	} else {
		c.sec -= minutes * 60
	}

	return c
}
