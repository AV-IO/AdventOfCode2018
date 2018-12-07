package main

import "testing"

func Test_ReadInstructions(t *testing.T) {
	tests := []struct {
		name            string
		instructions    []string
		wantsharedIn    int
		wantintactPatch int
	}{
		{"example 1-1", []string{"#1 @ 1,3: 4x4", "#2 @ 3,1: 4x4", "#3 @ 5,5: 2x2"}, 4, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotsharedIn, gotintactPatch := ReadInstructions(tt.instructions)
			if gotsharedIn != tt.wantsharedIn {
				t.Errorf("ReadInstructions() sharedIn: got %v, wanted %v", gotsharedIn, tt.wantsharedIn)
			}
			if gotintactPatch != tt.wantintactPatch {
				t.Errorf("ReadInstructions() intactPatch: got %v, wanted %v", gotintactPatch, tt.wantintactPatch)
			}
		})
	}
}
