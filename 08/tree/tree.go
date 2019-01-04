package tree

// Node is used for a node within a Tree
// Node consists of a list of pointers to children nodes, and a list of ints cointaining metadata
type Node struct {
	ID       int
	Children []int
	Metadata []int
}

// Tree is a map of IDs to nodes, and a pointer to the Root, and an internal counter for treeIDs
type Tree struct {
	RootID    int
	nodes     map[int]Node
	idCounter int
}

// NewTree : creates and initializes a new Tree
func NewTree() (t *Tree) {
	t = new(Tree)
	t.nodes = make(map[int]Node)
	return
}

// AddNode : Adds a node to the existing Tree under the ptr node with blank children and blank metadata.
func (t *Tree) AddNode(parentID int) (newID int) {
	newID = t.idCounter
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

	t.idCounter++
	return
}

// AddMetadata : Adds metadata to the specified node
func (t *Tree) AddMetadata(nodeID int, metadata []int) {
	t.nodes[nodeID] = Node{
		nodeID,
		t.nodes[nodeID].Children,
		metadata,
	}
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
