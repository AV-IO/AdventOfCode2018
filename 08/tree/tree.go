package tree

// Node is used for a node within a Tree
// Node consists of a list of pointers to children nodes, and a list of ints cointaining metadata
type Node struct {
	ID       int
	Children []int
	Metadata []int
}

// Tree is a list of nodes, and a pointer to the Root.
type Tree struct {
	nodes  map[int]Node
	RootID int
}

var idCounter int

// AddNode : Adds a node to the existing tree under the ptr node with blank children and blank metadata.
func (t *Tree) AddNode(parentID int) (newID int) {
	newID = idCounter
	t.nodes[newID] = Node{newID, []int{}, []int{}}

	if parentID != newID {
		t.nodes[parentID] = Node{ // Go: CaNnOt AsSiGn To StRuCt FiElD iN mAp
			parentID,
			append(t.nodes[parentID].Children, newID),
			t.nodes[parentID].Metadata,
		}
	} else {
		t.RootID = 0
	}

	idCounter++
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
