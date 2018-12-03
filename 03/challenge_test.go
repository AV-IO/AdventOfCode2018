package main

import "testing"

func Test_ReadInstructions(t *testing.T) {
	tests := []struct {
		name         string
		instructions []string
		wantsharedIn int
	}{
		{"example 1-1", []string{"#1 @ 1,3: 4x4", "#2 @ 3,1: 4x4", "#3 @ 5,5: 2x2"}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotsharedIn := ReadInstructions(tt.instructions); gotsharedIn != tt.wantsharedIn {
				t.Errorf("ReadInstructions() sharedIn: got %v, wanted %v", gotsharedIn, tt.wantsharedIn)
			}
		})
	}
}
