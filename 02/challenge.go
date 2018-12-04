package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

// FindChecksum :
//  	parameters:
//  		boxes: string slice, containing list of box IDs
//  		Channels:
//  			checksumChan: channel for sending back checksum of all boxes
func FindChecksum(boxes []string, checksumChan chan int) {
	doubles := 0
	triples := 0

	for _, ID := range boxes {
		seenDouble := false
		seenTriple := false
		for _, char := range ID {
			count := strings.Count(ID, string(char))
			if count == 2 {
				seenDouble = true
			} else if count == 3 {
				seenTriple = true
			}
		}
		if seenDouble {
			doubles++
		}
		if seenTriple {
			triples++
		}
	}

	checksumChan <- doubles * triples
}

// FindSimilar :
//  	parameters:
//  		boxes: string slice, containing list of box IDs
//  		Channels:
//  			similarChan: channel for sending back list of characters shared between the two most similar strings
func FindSimilar(boxes []string, similarChan chan string) {
	sort.Strings(boxes)

	for i := 0; i < len(boxes)-1; i++ {
		foundDiff := 0
		for j := i + 1; j < len(boxes); j++ {
			similar := boxes[i]
			for k := 0; k < len(boxes[i]); k++ {
				if boxes[i][k] != boxes[j][k] {
					if foundDiff > 0 {
						foundDiff++
						break
					}
					foundDiff++
					similar = similar[:k] + similar[k+1:]
				}
			}
			if foundDiff == 1 {
				similarChan <- similar
				return
			}
		}
	}
	similarChan <- ""
}

func main() {
	data, _ := ioutil.ReadFile("./input")
	checksum := make(chan int)
	similar := make(chan string)
	go FindChecksum(strings.Split(string(data), "\n"), checksum)
	go FindSimilar(strings.Split(string(data), "\n"), similar)
	output := "checksum: " + strconv.Itoa(<-checksum) + "\nsimilar characters: " + <-similar + "\n"
	fmt.Println(output)
	ioutil.WriteFile("./output", []byte(output), 0644)
}
