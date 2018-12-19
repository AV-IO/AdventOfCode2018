package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	s "strings"
)

type node struct {
	name string
	in   []*node
	out  []*node
}

type graph struct {
	nodes []node
	nmap  map[string]*node // Go: MaP vAlUeS aReN't AdDrEsSaBlE wItH pOiNtErS
}

func (g *graph) init() {
	g.nodes = make([]node, 0)
	g.nmap = make(map[string]*node)
}

func (g *graph) checkNode(name string) (exists bool) {
	_, exists = g.nmap[name]
	return
}

func (g *graph) addNode(name string) {
	g.nodes = append(g.nodes, node{name, make([]*node, 0), make([]*node, 0)})
	g.nmap[name] = &g.nodes[len(g.nodes)-1]
}

func (g *graph) addEdge(inName string, outName string) {
	g.nmap[inName].in = append(g.nmap[inName].in, g.nmap[outName])
	g.nmap[outName].out = append(g.nmap[outName].out, g.nmap[inName])
}

func (g *graph) instAdd(node string, requires string) {
	// Add nodes if thhey don't exist yet
	if !g.checkNode(node) {
		g.addNode(node)
	}
	if !g.checkNode(requires) {
		g.addNode(requires)
	}
	// Not going to check if edges exist

	g.addEdge(node, requires)
}

// readInstructions :
//  	parameters:
//  		instructions: list of instructions
func readInstructions(instructions []string) (instOrder string) {
	var g graph
	g.init()

	// Create graph
	for _, inst := range instructions {
		g.instAdd(string(inst[36]), string(inst[5]))
	}

	available := make([]string, 0) // List for available nodes
	// Find available
	for s, n := range g.nmap {
		if len(n.in) == 0 {
			available = append(available, s)
		}
	}

	for len(g.nmap) > 0 {
		sort.Strings(available) // making list alphabetical to get first alphabetical in available
		instOrder += available[0]

		for nidex, n := range g.nmap[available[0]].out { // Go: RaNgE iS oNlY bY vAlUe
			// delete incoming links from newly freed nodes
			if len(n.in) == 1 {
				g.nmap[available[0]].out[nidex].in = []*node{}
				// incoming node list is empty
				available = append(available, n.name)
			} else {
				for inNidex, inN := range n.in {
					if inN.name == available[0] {
						g.nmap[available[0]].out[nidex].in = append(
							g.nmap[available[0]].out[nidex].in[:inNidex],
							g.nmap[available[0]].out[nidex].in[inNidex+1:]...,
						)
						break
					}
				}
			}
		}

		// delete node from map, leaving slice since it's more process intensive to remove.
		delete(g.nmap, available[0])
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
