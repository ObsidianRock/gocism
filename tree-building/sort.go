package tree

import "sort"

func sorter(records []Record) []Record {

	// Closures that order the Change structure.
	increasingParent := func(c1, c2 *Record) bool {
		return c1.Parent < c2.Parent
	}

	increasingId := func(c1, c2 *Record) bool {
		return c1.ID < c2.ID
	}

	// Simple use: Sort by user.
	OrderedBy(increasingParent, increasingId).Sort(records)

	return records
}

type lessFunc func(p1, p2 *Record) bool

type multiSorter struct {
	records []Record
	less    []lessFunc
}

// Sort sorts the argument slice according to the less functions passed to OrderedBy.
func (ms *multiSorter) Sort(records []Record) {
	ms.records = records
	sort.Sort(ms)
}

// OrderedBy returns a Sorter that sorts using the less functions, in order.
// Call its Sort method to sort the data.
func OrderedBy(less ...lessFunc) *multiSorter {
	return &multiSorter{
		less: less,
	}
}

// Len is part of sort.Interface.
func (ms *multiSorter) Len() int {
	return len(ms.records)
}

// Swap is part of sort.Interface.
func (ms *multiSorter) Swap(i, j int) {
	ms.records[i], ms.records[j] = ms.records[j], ms.records[i]
}

func (ms *multiSorter) Less(i, j int) bool {

	p, q := &ms.records[i], &ms.records[j]
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
	// All comparisons to here said "equal", so just return whatever
	// the final comparison reports.
	return ms.less[k](p, q)
}
