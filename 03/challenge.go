package main

import (
	"fmt"
	"io/ioutil"
	sc "strconv"
	s "strings"
)

// ReadInstructions :
//  	parameters:
//  		instructions: instructions for requested squares of fabric
//		return values:
//  		sharedIn: the number of shared square inches
func ReadInstructions(instructions []string) (sharedIn int, lonePatch int) {
	var fabric [1000][1000]int
	fab := map[[2]int][]int{}
	fab[[2]int{2, 2}] = []int{4}

	for _, inst := range instructions {
		i := s.Split(inst, " ")
		xOff, _ := sc.Atoi(i[2][:s.Index(i[2], ",")])
		yOff, _ := sc.Atoi(i[2][s.Index(i[2], ",")+1 : len(i[2])-1])
		xSize, _ := sc.Atoi(i[3][:s.Index(i[3], "x")])
		ySize, _ := sc.Atoi(i[3][s.Index(i[3], "x")+1:])

		for x := 0; x < xSize; x++ {
			for y := 0; y < ySize; y++ {
				fabric[xOff+x][yOff+y]++
				if fabric[xOff+x][yOff+y] == 2 {
					sharedIn++
				}
			}
		}
	}

	return
}

func main() {
	data, _ := ioutil.ReadFile("./input")
	sharedIn := ReadInstructions(s.Split(string(data), "\r\n"))
	output := "shared inches: " + sc.Itoa(sharedIn) + "\n"
	fmt.Println(output)
	ioutil.WriteFile("./output", []byte(output), 0644)
}
