package sublist

type Relation string

//Sublist ...
func Sublist(list1, list2 []int) Relation {

	if len(list1) == len(list2) {
		for i := 0; i < len(list1); i++ {
			if list2[i] != list1[i] {
				return Relation("unequal")
			}
		}
		return Relation("equal")
	}

	if len(list1) > len(list2) {
		sub := sublist(list2, list1)
		if sub {
			return Relation("superlist")
		}
		return Relation("unequal")
	}

	sub := sublist(list1, list2)
	if sub {
		return Relation("sublist")
	}

	return Relation("unequal")
}

func sublist(list1, list2 []int) bool {

	if len(list1) == 0 {
		return true
	}

	end, front := 0, len(list1)-1

	for front < len(list2) {

		var count int
		for i, j := 0, end; j <= front; i, j = i+1, j+1 {

			if list2[j] == list1[i] {
				count++
			}
			if count == len(list1) {
				return true
			}

		}

		front++
		end++

	}
	return false
}
