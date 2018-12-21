package tree

import (
	"fmt"
)

// Record hold the structure for a record
type Record struct {
	ID, Parent int
}

// Node holds the structure for a Node
type Node struct {
	ID       int
	Children []*Node
}

// Mismatch holds structore for Mismatch
type Mismatch struct{}

func (m Mismatch) Error() string {
	return "c"
}

// Build builds the records
func Build(records []Record) (*Node, error) {

	if len(records) < 1 {
		return nil, nil
	}

	// mapping := make(map[int][]int)

	root := &Node{}

	for k, v := range records {

		fmt.Printf("v is %#v\n", v)
		//		fmt.Printf("ID is: %+v and for ID %+v, Parent is %+v\n", v.ID, v.ID, v.Parent)
		fmt.Printf("mapping[%v][%v] = %v\n", v.Parent, k, v.ID)
	}
	/* for i := range mapping {
		fmt.Println(i)
	} */

	return root, nil
}
