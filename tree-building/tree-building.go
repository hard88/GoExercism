package tree

import (
	"errors"
	"sort"
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

type ByID []Record

func (records ByID) Len() int           { return len(records) }
func (records ByID) Swap(i, j int)      { records[i], records[j] = records[j], records[i] }
func (records ByID) Less(i, j int) bool { return records[i].ID < records[j].ID }

func Build(records []Record) (*Node, error) {

	if len(records) == 0 {
		return nil, nil
	}

	root := Node{}

	sort.Sort(ByID(records))

	for i := range records {

		if i == 0 {
			if records[i].ID != records[i].Parent {
				return nil, errors.New("no root node")
			}
			root.ID = records[i].ID
		}

		if records[i].Parent > 0 && records[i].ID <= records[i].Parent {
			return nil, errors.New("children id bigger than parent")
		}

		if records[i].ID >= len(records) {
			return nil, errors.New("out of range")
		}

	}

	for i := range records {
		if i == 0 {
			continue
		}

		node := Node{ID: records[i].ID}
		isBelong := false

		if records[i].ID == records[i-1].ID {
			return nil, errors.New("duplicate node")
		}

		if records[i].Parent == root.ID {
			root.Children = append(root.Children, &node)
			isBelong = true
		} else {
			for _, c := range root.Children {
				if c.ID == records[i].Parent {
					c.Children = append(c.Children, &node)
					isBelong = true
					break
				}
			}
		}

		if !isBelong {
			return nil, errors.New("no parent")
		}
	}

	return &root, nil
}
