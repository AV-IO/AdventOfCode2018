package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// ReadInstruction reads from a list of instructions and outputs two frequencies
// parameters:
//     instructions: string slice, containing list of instructions
// return values:
//     f1: end frequency
//     f2: first duplicate frequency seen
func ReadInstruction(instructions []string) (f1 int, f2 int) {
	current := 0
	seen := map[int]bool{0: true}
	f1set := false
	f2set := false

	for lcount := 0; f2set == false; lcount++ {
		for i, inst := range instructions {
			ii, _ := strconv.Atoi(inst)
			current += ii
			if f1set == false && i == len(instructions)-1 {
				f1 = current
				f1set = true
			}
			if seen[current] == false {
				seen[current] = true
			} else {
				f2 = current
				f2set = true
				break
			}
		}
		if lcount == 100000 { //stop infinite loop
			f2 = 999999
			f2set = true
			break
		}
	}

	return
}

func main() {
	data, _ := ioutil.ReadFile("./input")
	f1, f2 := ReadInstruction(strings.Split(string(data), "\n"))
	output := "frequency 1: " + strconv.Itoa(f1) + "\nfrequency 2: " + strconv.Itoa(f2) + "\r\n"
	fmt.Println(output)
	ioutil.WriteFile("./output", []byte(output), 0644)
}
