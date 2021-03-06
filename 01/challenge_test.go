package main

import (
	"testing"
)

func TestReadInstruction(t *testing.T) {
	type parameters struct {
		args []string
	}
	tests := []struct {
		name   string
		param  parameters
		wantF1 int
		wantF2 int
	}{
		{"example 1-1", parameters{[]string{"+1", "-2", "+3", "+1"}}, 3, 0},
		{"example 1-2", parameters{[]string{"+1", "+1", "+1"}}, 3, 0},
		{"example 1-3", parameters{[]string{"+1", "+1", "-2"}}, 0, 0},
		{"example 1-4", parameters{[]string{"-1", "-2", "-3"}}, -6, 0},

		{"example 2-1", parameters{[]string{"+1", "-2", "+3", "+1"}}, 3, 2},
		{"example 2-2", parameters{[]string{"+1", "-1"}}, 0, 0},
		{"example 2-3", parameters{[]string{"+3", "+3", "+4", "-2", "-4"}}, 4, 10},
		{"example 2-4", parameters{[]string{"-6", "+3", "+8", "+5", "-6"}}, 4, 5},
		{"example 2-5", parameters{[]string{"+7", "+7", "-2", "-7", "-4"}}, 1, 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotF1, gotF2 := ReadInstruction(tt.param.args)
			if gotF1 != tt.wantF1 {
				t.Errorf("ReadInstruction() f1: got %v, wanted %v", gotF1, tt.wantF1)
			}
			if tt.name[8] == []byte("2")[0] && gotF2 != tt.wantF2 {
				t.Errorf("ReadInstruction() f2: got %v, wanted %v", gotF2, tt.wantF2)
			}
		})
	}
}
