package main

import "testing"

func Test_readInstructions(t *testing.T) {
	type parameters struct {
		instructions []string
	}
	tests := []struct {
		name          string
		param         parameters
		wantInstOrder string
	}{
		{
			"example 1-1",
			parameters{[]string{
				"Step C must be finished before step A can begin.",
				"Step C must be finished before step F can begin.",
				"Step A must be finished before step B can begin.",
				"Step A must be finished before step D can begin.",
				"Step B must be finished before step E can begin.",
				"Step D must be finished before step E can begin.",
				"Step F must be finished before step E can begin.",
			}},
			"CABDFE",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotInstOrder := readInstructions(tt.param.instructions); gotInstOrder != tt.wantInstOrder {
				t.Errorf("readInstructions() instOrder: got %v, wanted %v", gotInstOrder, tt.wantInstOrder)
			}
		})
	}
}
