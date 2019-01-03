package main

import (
	"../08/tree"
	"fmt"
	"io/ioutil"
	sc "strconv"
	s "strings"
)

//	readTree :
//  	parameters:
//  		treeData: list of ints continating instructions for forming the tree
//		return values:
//			sum: sum of all metadata values in tree
func readTree(treeData []int) (sum int) {
	var t tree.Tree
	ptrStack := []*tree.Node{}
	instructionReadRec(&t, ptrStack, treeData, 0)
	t.PrintConnectionCount()
	sum = t.MetadataSum()
	return
}

//	instructionReadRec :
//		parameters:
//			t: tree
//			ptrStack: pointer stack for tracking location in tree
//			treeData: list of data for importing into tree
//			treeDataIndex: current location in treeData
//		return values:
//			newTreeDataIndex: updated location in treeData
func instructionReadRec(
	t *tree.Tree,
	ptrStack []*tree.Node,
	treeData []int,
	treeDataIndex int,
) (newTreeDataIndex int) {
	childrenCount := treeData[treeDataIndex]
	metadataCount := treeData[treeDataIndex+1]
	treeDataIndex += 2

	// Add new node, and add node pointer to ptrStack
	var newPtr *tree.Node
	if len(ptrStack) == 0 {
		newPtr = t.AddNode(nil)
	} else {
		newPtr = t.AddNode(ptrStack[len(ptrStack)-1])
	}
	ptrStack = append(ptrStack, newPtr)

	for i := 0; i < childrenCount; i++ { // For each child, call recursive
		treeDataIndex = instructionReadRec(t, ptrStack, treeData, treeDataIndex)
	}
	for i := 0; i < metadataCount; i++ { // For each metadata, add to metadata list.
		newPtr.Metadata = append(newPtr.Metadata, treeData[treeDataIndex])
		treeDataIndex++
	}

	return treeDataIndex
}

func main() {
	data, _ := ioutil.ReadFile("./input")
	treeData := []int{}
	for _, n := range s.Split(s.Trim(string(data), "\n"), " ") {
		nint, _ := sc.Atoi(n)
		treeData = append(treeData, nint)
	}

	sum := readTree(treeData)
	output := "sum of treeData: " + sc.Itoa(sum) + "\n"
	fmt.Println(output)
	ioutil.WriteFile("./output", []byte(output), 0644)
}
