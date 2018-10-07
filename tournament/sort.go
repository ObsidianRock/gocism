package tournament

import "sort"

type lessFunc func(p1, p2 *Team) bool

type multiSorter struct {
	teams []Team
	less  []lessFunc
}

func (ms *multiSorter) Sort(teams []Team) {
	ms.teams = teams
	sort.Sort(ms)
}

func orderedBy(less ...lessFunc) *multiSorter {
	return &multiSorter{
		less: less,
	}
}

func (ms *multiSorter) Len() int {
	return len(ms.teams)
}

func (ms *multiSorter) Swap(i, j int) {
	ms.teams[i], ms.teams[j] = ms.teams[j], ms.teams[i]
}

func (ms *multiSorter) Less(i, j int) bool {
	p, q := &ms.teams[i], &ms.teams[j]
	// Try all but the last comparison.
	var k int
	for k = 0; k < len(ms.less)-1; k++ {
		less := ms.less[k]
		switch {
		case less(p, q):
			// p < q, so we have a decision.
			return true
		case less(q, p):
			// p > q, so we have a decision.
			return false
		}
		// p == q; try the next comparison.
	}
	return ms.less[k](p, q)
}

func sortTeam(teams []Team) {

	// Closures that order the Change structure.
	name := func(c1, c2 *Team) bool {
		return c1.Name < c2.Name
	}

	score := func(c1, c2 *Team) bool {
		return c1.Total > c2.Total
	}

	orderedBy(score, name).Sort(teams)
}
