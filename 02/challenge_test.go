package main

import (
	"testing"
)

func TestFindChecksum(t *testing.T) {
	type parameters struct {
		boxes        []string
		checksumChan chan int
	}
	checksum := make(chan int)
	tests := []struct {
		name         string
		param        parameters
		wantChecksum int
	}{
		{"example 1-1", parameters{[]string{"abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"}, checksum}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go FindChecksum(tt.param.boxes, tt.param.checksumChan)
			if gotChecksum := <-checksum; gotChecksum != tt.wantChecksum {
				t.Errorf("ReadIDs() checksum: got %v, wanted %v", gotChecksum, tt.wantChecksum)
			}
		})
	}
}

func TestFindSimilar(t *testing.T) {
	type parameters struct {
		boxes       []string
		similarChan chan string
	}
	similar := make(chan string)
	tests := []struct {
		name        string
		param       parameters
		wantSimilar string
	}{
		{"example 2-1", parameters{[]string{"abcde", "fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"}, similar}, "fgij"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go FindSimilar(tt.param.boxes, tt.param.similarChan)
			if gotSimilar := <-similar; gotSimilar != tt.wantSimilar {
				t.Errorf("FindSimilar() similar: got %v, wanted %v", gotSimilar, tt.wantSimilar)
			}
		})
	}
}
