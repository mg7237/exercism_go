package tree

import (
	"errors"
	"sort"
)

//Record struct defines structure of input records
type Record struct {
	ID     int
	Parent int
}

// Node defines structure of a Node
type Node struct {
	ID       int
	Children []*Node
}

// Build Nodes
func Build(records []Record) (*Node, error) {

	m := make(map[int]*Node)

	sort.Slice(records, func(i, j int) bool { return records[i].ID < records[j].ID })

	for i, r := range records {
		if r.ID != i || r.Parent > r.ID || (r.ID > 0 && r.Parent == r.ID) {
			return nil, errors.New("not in sequence or has bad parent")
		}

		m[i] = &Node{ID: i}

		if i != 0 {
			m[r.Parent].Children = append(m[r.Parent].Children, m[i])
		}

	}
	return m[0], nil
}
