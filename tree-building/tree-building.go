package tree

import (
	"fmt"
)

type Record struct {
	ID       int
	Parent   int
	Children []*Node
}

type Node struct {
	ID       int
	Children []*Node
}

const rootID = 0

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	for i, r := range records {
		if r.ID < rootID || r.ID >= len(records) {
			return nil, fmt.Errorf("out of bounds of id %d", r.ID)
		}
		if r.ID != i {
			records[i], records[r.ID] = records[r.ID], records[i]
		}
	}

	nodes := make([]Node, len(records))

	for i, r := range records {
		if r.ID != i {
			return nil, fmt.Errorf("non-contiguous node, want %d, get %d", i, r.ID)
		}

		validParentForChild := (r.ID > r.Parent) || (r.ID == rootID && r.Parent == rootID)
		if !validParentForChild {
			return nil, fmt.Errorf("node %d has impossible parent %d", r.ID, r.Parent)
		}

		nodes[i].ID = i
		if i != rootID {
			q := &nodes[r.Parent]
			q.Children = append(q.Children, &nodes[i])
		}

	}

	return &nodes[0], nil

}
