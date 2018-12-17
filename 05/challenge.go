package main

import (
	"fmt"
	"io/ioutil"
	sc "strconv"
	s "strings"
)

// AlchemicalReduction :
//  	parameters:
//  		inPolymer: input polymer string
//  	channels:
//  		outPolyLen: returns length of ouput polymer after being reduced
func AlchemicalReduction(inPolymer string, outPolyLen chan int) {
	// make replace list
	var repList []string
	for i := 0; i < 26; i++ {
		repList = append(repList, string(97+i)+string(65+i)) //aA
		repList = append(repList, string(65+i)+string(97+i)) //Aa
	}

	oldLen := len(inPolymer) + 1
	outPolymer := inPolymer
	// while the length of the polymer is still decreasing
	for {
		// not incredibly efficient, but checking which characters are still in outPolymer may be less efficient...
		for _, rep := range repList {
			outPolymer = s.Replace(outPolymer, rep, "", -1)
		}
		if len(outPolymer) < oldLen {
			oldLen = len(outPolymer)
		} else {
			break
		}
	}

	outPolyLen <- len(outPolymer)
}

// FindProblemUnit :
//  	parameters:
//  		inPolymer: input polymer string
//  	channels:
//  		ProblemUnit: returns problem unit in polymer
//			outFixedPolyLen: returns length of polymer after removing problem unit
func FindProblemUnit(inPolymer string, ProblemUnit chan string, outFixedPolyLen chan int) {
	polyLenTest := make(chan int)

	var unit string
	len := len(inPolymer)

	for i := 0; i < 26; i++ {
		outPolymer := s.Replace(inPolymer, string(97+i), "", -1)
		outPolymer = s.Replace(outPolymer, string(65+i), "", -1)

		go AlchemicalReduction(outPolymer, polyLenTest)

		lenTest := <-polyLenTest
		if lenTest < len {
			len = lenTest
			unit = string(97+i) + string(65+i)
		}
	}

	ProblemUnit <- unit
	outFixedPolyLen <- len
}

func main() {
	data, _ := ioutil.ReadFile("./input")
	outPolyLen := make(chan int)
	ProblemUnit := make(chan string)
	outFixedPolyLen := make(chan int)
	go AlchemicalReduction(s.Split(string(data), "\n")[0], outPolyLen)
	go FindProblemUnit(s.Split(string(data), "\n")[0], ProblemUnit, outFixedPolyLen)
	output := "Output polymer length: " + sc.Itoa(<-outPolyLen) +
		"\nProblem Unit: " + <-ProblemUnit +
		"\nLength after removing unit: " + sc.Itoa(<-outFixedPolyLen)
	fmt.Println(output)
	ioutil.WriteFile("./output", []byte(output), 0644)
}
