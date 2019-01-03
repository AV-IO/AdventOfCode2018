package tree

import (
	"fmt"
)

// Node is used for a node within a Tree
// Node consists of a list of pointers to children nodes, and a list of ints cointaining metadata
type Node struct {
	Children []*Node
	Metadata []int
}

// Tree is a list of nodes, and a pointer to the Root.
type Tree struct {
	nodes []Node
	Root  *Node
}

// AddNode : Adds a node to the existing tree under the ptr node with blank children and blank metadata.
func (t *Tree) AddNode(ptr *Node) (newPtr *Node) {
	// TODO: check if default initialization is ok to append to.
	t.nodes = append(
		t.nodes,
		Node{[]*Node{}, []int{}},
	)
	newPtr = &t.nodes[len(t.nodes)-1]

	if ptr != nil {
		ptr.Children = append(ptr.Children, newPtr)
	} else {
		t.Root = newPtr
	}
	return
}

// MetadataSum : gets total sum of all metadata in tree
func (t *Tree) MetadataSum() (sum int) {
	for _, n := range t.nodes {
		for _, i := range n.Metadata {
			sum += i
		}
	}
	return
}

// PrintConnectionCount : prints list of all child connections
func (t *Tree) PrintConnectionCount() {
	for _, n := range t.nodes {
		fmt.Println(len(n.Children))
	}
}
