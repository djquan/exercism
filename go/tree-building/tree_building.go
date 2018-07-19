/*Package tree provides functionality to convert an unstructured
slice of records into a structured tree node
*/
package tree

import (
	"fmt"
	"sort"
)

// Record represents an unstructured tree node
type Record struct {
	ID, Parent int
}

// Node represents a tree node
type Node struct {
	ID       int
	Children []*Node
}

// Build converts a slice of records to a Tree
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	nodes := make([]*Node, len(records))

	for i, r := range records {
		nodes[i] = &Node{ID: r.ID}

		if i == 0 && (r.Parent != 0 || r.ID != 0) {
			return nil, fmt.Errorf("Invalid root record")
		} else if i == 0 {
			continue
		} else if i != r.ID || r.ID <= r.Parent {
			return nil, fmt.Errorf("Invalid input")
		}

		if parent := nodes[r.Parent]; parent != nil {
			parent.Children = append(parent.Children, nodes[i])
		}
	}

	return nodes[0], nil
}
