package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(f func(int) bool) Ints {

	var ints Ints

	for _, v := range i {
		if f(v) {
			ints = append(ints, v)
		}
	}
	return ints
}

func (i Ints) Discard(f func(int) bool) Ints {

	var ints Ints
	for _, v := range i {
		if !f(v) {
			ints = append(ints, v)

		}
	}
	return ints

}

func (l Lists) Keep(f func([]int) bool) Lists {

	var lists Lists

	for _, v := range l {
		if f(v) {
			lists = append(lists, v)
		}

	}
	return lists

}

func (s Strings) Keep(f func(string) bool) Strings {

	var str Strings

	for _, v := range s {
		if f(v) {
			str = append(str, v)
		}
	}
	return str
}
