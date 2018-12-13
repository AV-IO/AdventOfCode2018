package main

import (
	"testing"
)

func TestAlchemicalReduction(t *testing.T) {
	type parameters struct {
		inPolymer  string
		outPolyLen chan int
	}
	tests := []struct {
		name        string
		param       parameters
		wantPolyLen int
	}{
		{"example 1-1", parameters{"aA", make(chan int)}, 0},
		{"example 1-2", parameters{"abBA", make(chan int)}, 0},
		{"example 1-3", parameters{"abAB", make(chan int)}, 4},
		{"example 1-4", parameters{"aabAAB", make(chan int)}, 6},
		{"example 1-5", parameters{"dabAcCaCBAcCcaDA", make(chan int)}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AlchemicalReduction(tt.param.inPolymer, tt.param.outPolyLen)
			if gotPolyLen := <-tt.param.outPolyLen; gotPolyLen != tt.wantPolyLen {
				t.Errorf(
					"AlchemicalReduction() outPolyLen: got %v, wanted %v)",
					gotPolyLen,
					tt.wantPolyLen,
				)
			}
		})
	}
}

func TestFindProblemUnit(t *testing.T) {
	type parameters struct {
		inPolymer       string
		ProblemUnit     chan string
		outFixedPolyLen chan int
	}
	tests := []struct {
		name        string
		param       parameters
		wantProblem string
		wantPolyLen int
	}{
		{"example 2-1", parameters{"dabAcCaCBAcCcaDA", make(chan string), make(chan int)}, "cC", 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FindProblemUnit(tt.param.inPolymer, tt.param.ProblemUnit, tt.param.outFixedPolyLen)
			gotProblem := <-tt.param.ProblemUnit
			gotPolyLen := <-tt.param.outFixedPolyLen
			if gotProblem != tt.wantProblem {
				t.Errorf(
					"FindProblemUnit() ProblemUnit: got %v, wanted %v",
					gotProblem,
					tt.wantProblem,
				)
			}
			if gotPolyLen != tt.wantPolyLen {
				t.Errorf(
					"FindProblemUnit() outFixedPolyLen: got %v, wanted %v",
					gotPolyLen,
					tt.wantPolyLen,
				)
			}
		})
	}
}
