package main

import (
	"../07/graph"
	"fmt"
	"io/ioutil"
	"sort"
	s "strings"
)

// readInstructions :
//  	parameters:
//  		instructions: list of instructions
func readInstructions(instructions []string) (instOrder string) {
	var g graph.Graph
	g.Init()

	// Create graph
	for _, inst := range instructions {
		g.InstAdd(string(inst[36]), string(inst[5]))
	}

	available := make([]string, 0) // List for available nodes
	// Find available
	for s, n := range g.Nmap {
		if len(n.In) == 0 {
			available = append(available, s)
		}
	}

	for len(g.Nmap) > 0 {
		sort.Strings(available) // making list alphabetical to get first alphabetical in available
		instOrder += available[0]

		for nidex, n := range g.Nmap[available[0]].Out { // Go: RaNgE iS oNlY bY vAlUe
			// delete incoming links from newly freed nodes
			if len(n.In) == 1 {
				g.Nmap[available[0]].Out[nidex].In = []*graph.Node{}
				// incoming node list is empty
				available = append(available, n.Name)
			} else {
				for inNidex, inN := range n.In {
					if inN.Name == available[0] {
						g.Nmap[available[0]].Out[nidex].In = append(
							g.Nmap[available[0]].Out[nidex].In[:inNidex],
							g.Nmap[available[0]].Out[nidex].In[inNidex+1:]...,
						)
						break
					}
				}
			}
		}

		// delete node from map, leaving slice since it's more process intensive to remove.
		delete(g.Nmap, available[0])
		available = available[1:]
	}

	return
}

func main() {
	data, _ := ioutil.ReadFile("./input")
	instOrder := readInstructions(s.Split(string(data), "\n"))
	output := "Order of instructions: " + instOrder + "\n"
	fmt.Println(output)
	ioutil.WriteFile("./output", []byte(output), 0644)
}
