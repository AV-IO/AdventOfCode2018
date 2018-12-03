package main

import (
	"testing"
)

func TestFindChecksum(t *testing.T) {
	tests := []struct {
		name         string
		boxes        []string
		wantChecksum int
	}{
		{"example 1-1", []string{"abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotChecksum := FindChecksum(tt.boxes)
			if gotChecksum != tt.wantChecksum {
				t.Errorf("ReadIDs() checksum: got %v, wanted %v", gotChecksum, tt.wantChecksum)
			}
		})
	}
}

func TestFindSimilar(t *testing.T) {
	tests := []struct {
		name        string
		boxes       []string
		wantSimilar string
	}{
		{"example 2-1", []string{"abcde", "fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"}, "fgij"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSimilar := FindSimilar(tt.boxes); gotSimilar != tt.wantSimilar {
				t.Errorf("FindSimilar() similar: got %v, wanted %v", gotSimilar, tt.wantSimilar)
			}
		})
	}
}
