package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

//Robot name
type Robot string

var rnd *rand.Rand

func init() {
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
}

//New ...
func New() *Robot {
	return new(Robot)
}

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

//Name ...
func (r *Robot) Name() string {

	if len(*r) == 0 {
		*r = Robot(nameGen())
		return string(*r)
	}

	return string(*r)
}

// Reset ...
func (r *Robot) Reset() {

	*r = Robot(nameGen())
}

// range specification, note that min <= max
type intRange struct {
	min, max int
}

// get next random value within the interval including min and max
func (ir *intRange) nextRandom(r *rand.Rand) int {
	return r.Intn(ir.max-ir.min+1) + ir.min
}

func numberGen(rnd *rand.Rand) string {
	ir := intRange{0, 999}
	return fmt.Sprintf("%03d", ir.nextRandom(rnd))
}

func letterGen(rnd *rand.Rand) string {
	ir := intRange{0, len(charset) - 1}

	var b string

	for i := 0; i < 2; i++ {
		t := ir.nextRandom(rnd)
		v := string(charset[t])
		b += v

	}
	return b
}

func nameGen() string {

	letter := letterGen(rnd)
	num := numberGen(rnd)

	name := letter + num

	return name
}
