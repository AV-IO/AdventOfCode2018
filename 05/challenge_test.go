package main

import "testing"

func TestAlchemicalReduction(t *testing.T) {
	tests := []struct {
		name              string
		inPolymer         string
		wantOutPolymerStr string
		wantOutPolymerLen int
	}{
		{"example 1-1", "aA", "", 0},
		{"example 1-2", "abBA", "", 0},
		{"example 1-3", "abAB", "abAB", 4},
		{"example 1-4", "aabAAB", "aabAAB", 6},
		{"example 1-5", "dabAcCaCBAcCcaDA", "dabCBAcaDA", 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: this is old. needs to be updated to match channel usage
			gotOutPolymer := AlchemicalReduction(tt.inPolymer)
			gotOutPolymerLen := len(gotOutPolymer)
			if gotOutPolymer != tt.wantOutPolymerStr {
				t.Errorf(
					"AlchemicalReduction() outPolymer: got %v (len: %v), wanted %v (len: %v)",
					gotOutPolymer,
					gotOutPolymerLen,
					tt.wantOutPolymerStr,
					tt.wantOutPolymerLen,
				)
			}
		})
	}
}
