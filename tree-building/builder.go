package tree

import (
	"fmt"
)

func Builder(records []Record) (*Node, error) {

	if len(records) == 0 {
		return nil, fmt.Errorf("empty records")
	}

	r := sorter(records)

	if err := detectErr(r); err != nil {
		return nil, fmt.Errorf("err detected with records: %s", err)
	}

	return maketree(r)
}

func detectErr(r []Record) error {

	if r[0].ID != 0 || r[0].Parent != 0 {
		return fmt.Errorf("parent node not consistent")
	}

	for i := 1; i < len(r); i++ {

		if r[i].ID == 0 {
			return fmt.Errorf("dublicate root")
		}

		if r[i].ID == r[i-1].ID {
			return fmt.Errorf("dublicate node")
		}

		if r[i].ID < r[i].Parent {
			return fmt.Errorf("higher id parent of lower id")
		}

		if r[i].ID == r[i].Parent {
			return fmt.Errorf("cycle directly")
		}
	}

	return nil
}

func maketree(records []Record) (*Node, error) {

	tree := &Node{}

	root := &Node{ID: records[0].ID}

	for i := 1; i < len(records); i++ {
		addNode(root, records[i])
	}
	return root, nil
}

func addNode(root *Node, r Record) {

	node := &Node{ID: r.ID}

	for {

	}

}
