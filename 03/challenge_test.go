package main

import "testing"

func Test_funcName(t *testing.T) {
	tests := []struct {
		name     string
		paramA   []string
		wantRetA int
	}{
		// Test Cases
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRetA := funcName(tt.paramA); gotRetA != tt.wantRetA {
				t.Errorf("funcName() = %v, want %v", gotRetA, tt.wantRetA)
			}
		})
	}
}
