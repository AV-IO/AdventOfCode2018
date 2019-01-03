package main

import "testing"

func Test_readTree(t *testing.T) {
	type parameters struct {
		treeData []int
	}
	tests := []struct {
		name    string
		param   parameters
		wantSum int
	}{
		{"example 1-1", parameters{[]int{2, 3, 0, 3, 10, 11, 12, 1, 1, 0, 1, 99, 2, 1, 1, 2}}, 138},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := readTree(tt.param.treeData); gotSum != tt.wantSum {
				t.Errorf("readTree() readTree: got %v, wanted %v", gotSum, tt.wantSum)
			}
		})
	}
}
