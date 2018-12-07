package main

import (
	"fmt"
	"io/ioutil"
	sc "strconv"
	s "strings"
)

// ReadInstructions :
//		parameters:
//			instructions: instructions for requested squares of fabric
//		return values:
//			sharedIn: the number of shared square inches
//			intactPatch:
func ReadInstructions(instructions []string) (sharedIn int, intactPatch int) {
	fabric := map[[2]int][]int{}
	intactPatches := make(map[int]bool)
	for i := 0; i < len(instructions); i++ {
		intactPatches[i+1] = false
	}

	for instN, inst := range instructions {
		i := s.Split(inst, " ")
		xOff, _ := sc.Atoi(i[2][:s.Index(i[2], ",")])
		yOff, _ := sc.Atoi(i[2][s.Index(i[2], ",")+1 : len(i[2])-1])
		xSize, _ := sc.Atoi(i[3][:s.Index(i[3], "x")])
		ySize, _ := sc.Atoi(i[3][s.Index(i[3], "x")+1:])

		for x := 0; x < xSize; x++ {
			for y := 0; y < ySize; y++ {
				coord := [2]int{xOff + x, yOff + y} // Go: MaP vAlUeS aReN't AdDrEsSaBlE wItH pOiNtErS
				fabric[coord] = append(fabric[coord], instN+1)
				if len(fabric[coord]) == 2 {
					sharedIn++
					delete(intactPatches, fabric[coord][0])
					delete(intactPatches, fabric[coord][1])
				} else if len(fabric[coord]) > 2 {
					delete(intactPatches, fabric[coord][len(fabric[coord])-1])
				}
			}
		}
	}

	for key := range intactPatches {
		intactPatch = key // just getting the first value
		break
	}

	return
}

func main() {
	data, _ := ioutil.ReadFile("./input")
	sharedIn, intactPatch := ReadInstructions(s.Split(string(data), "\r\n"))
	output := "shared inches: " + sc.Itoa(sharedIn) + "\nintact Patch: " + sc.Itoa(intactPatch) + "\n"
	fmt.Println(output)
	ioutil.WriteFile("./output", []byte(output), 0644)
}
