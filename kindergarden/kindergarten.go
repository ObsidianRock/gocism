package kindergarten

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

//Garden ...
type Garden struct {
	childrens []child
}

type child struct {
	name   string
	plants []string
}

/*
	grow grass, clover, radishes, and violets.

    Alice-0, Bob-1, Charlie-2, David-3,
    Eve-4, Fred-5, Ginny-6, Harriet-6,
    Ileana-7, Joseph-8, Kincaid-9, and Larry-10.

*/

var translate = map[string]string{"V": "violets", "G": "grass", "R": "radishes", "C": "clover"}

func (g *Garden) Plants(child string) ([]string, bool) {

	for _, v := range g.childrens {

		if v.name == child {
			return v.plants, true
		}
	}

	return nil, false
}

func check(lines string, children []string) (map[int][]string, error) {

	sort.Strings(children)

	for i := 0; i < len(children); i++ {

		if i+1 >= len(children) {
			break
		}

		next := children[i+1]
		if children[i] == next {
			return nil, fmt.Errorf("dublicate name")
		}
	}

	if lines[0] != '\n' {
		return nil, fmt.Errorf("wrong format")
	}

	s := strings.Split(lines[1:], "\n")

	prevLine := len(s[0])
	for _, v := range s {
		if len(v)%2 != 0 {
			return nil, fmt.Errorf("odd number of cups")
		}

		if len(v) != prevLine {
			return nil, fmt.Errorf("mismatched rows")
		}

	}

	x := make(map[int][]string)
	var count, index int
	for _, v := range lines {
		if unicode.IsSpace(v) {
			count, index = 0, 0
			continue
		}

		if v == 'C' || v == 'G' || v == 'R' || v == 'V' {
			if count%2 == 0 && count != 0 {
				count = 0
				index++
			}
			plant := translate[string(v)]

			x[index] = append(x[index], plant)
			count++
			continue
		}
		return nil, fmt.Errorf("invaid cup codes")
	}

	return x, nil
}

func NewGarden(diagram string, children []string) (*Garden, error) {

	var g Garden
	cpychld := make([]string, len(children))

	copy(cpychld, children)

	x, err := check(diagram, cpychld)
	if err != nil {
		return nil, fmt.Errorf("diagram: %s", err.Error())
	}

	for k, v := range x {
		var c child
		c.name = cpychld[k]
		c.plants = v
		g.childrens = append(g.childrens, c)
	}

	return &g, nil
}
