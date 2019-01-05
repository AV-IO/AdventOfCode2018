package tree

import (
	"testing"
)

func TestTree_MetadataSum(t *testing.T) {
	tests := []struct {
		name    string
		fields  Tree
		wantSum int
	}{
		{
			"Single Node",
			Tree{
				RootID: 0,
				nodes: map[int]Node{
					0: Node{ID: 0, Children: []int{}, Metadata: []int{1, 1, 2}},
				},
				idCounter: 1,
			},
			4,
		},
		{
			"Test Tree",
			Tree{
				RootID: 0,
				nodes: map[int]Node{
					0: Node{ID: 0, Children: []int{1, 2}, Metadata: []int{1, 1, 2}},
					1: Node{ID: 1, Children: []int{}, Metadata: []int{10, 11, 12}},
					2: Node{ID: 2, Children: []int{3}, Metadata: []int{2}},
					3: Node{ID: 3, Children: []int{}, Metadata: []int{99}},
				},
				idCounter: 4,
			},
			138,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Tree{
				RootID:    tt.fields.RootID,
				nodes:     tt.fields.nodes,
				idCounter: tt.fields.idCounter,
			}
			if gotSum := tr.MetadataSum(); gotSum != tt.wantSum {
				t.Errorf("Tree.MetadataSum() val: got %v, wanted %v", gotSum, tt.wantSum)
			}
		})
	}
}

func TestTree_NodeValue(t *testing.T) {
	type parameters struct {
		nodeID int
	}
	tr1 := &Tree{
		RootID: 0,
		nodes: map[int]Node{
			0: Node{ID: 0, Children: []int{1, 2}, Metadata: []int{1, 1, 2}},
			1: Node{ID: 1, Children: []int{}, Metadata: []int{10, 11, 12}},
			2: Node{ID: 2, Children: []int{3}, Metadata: []int{2}},
			3: Node{ID: 3, Children: []int{}, Metadata: []int{99}},
		},
		idCounter: 4,
	}
	tr2 := &Tree{
		RootID: 0,
		nodes: map[int]Node{
			0: Node{ID: 0, Children: []int{1, 2, 4}, Metadata: []int{1, 1, 2}},
			1: Node{ID: 1, Children: []int{}, Metadata: []int{1, 2, 3}},
			2: Node{ID: 2, Children: []int{3}, Metadata: []int{2}},
			3: Node{ID: 3, Children: []int{}, Metadata: []int{99}},
			4: Node{ID: 4, Children: []int{5}, Metadata: []int{1}},
			5: Node{ID: 5, Children: []int{}, Metadata: []int{10}},
		},
		idCounter: 4,
	}
	tests := []struct {
		name    string
		fields  *Tree
		param   parameters
		wantVal int
	}{
		{"Example Tree: A", tr1, parameters{0}, 66},
		{"Example Tree: C", tr1, parameters{2}, 0},
		{"Extended Tree: A", tr2, parameters{0}, 16},
		{"Extended Tree: E", tr2, parameters{4}, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotVal := tt.fields.NodeValue(tt.param.nodeID); gotVal != tt.wantVal {
				t.Errorf("Tree.NodeValue() val: got %v, wanted %v", gotVal, tt.wantVal)
			}
		})
	}
}
