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
func readTree(treeData []int) (sum int, val int) {
	t := tree.NewTree()
	instructionReadRec(t, 0, treeData, 0)
	sum = t.MetadataSum()
	val = t.NodeValue(t.RootID)
	return
}

//	instructionReadRec :
//		parameters:
//			t: tree
//			parentID: ID of the parent node which called the function
//			treeData: list of data for importing into tree
//				only slice header is passed by copy, so not space-intensive
//			treeDataIndex: current location in treeData
//		return values:
//			newTreeDataIndex: updated location in treeData
func instructionReadRec(
	t *tree.Tree,
	parentID int,
	treeData []int,
	treeDataIndex int,
) (newTreeDataIndex int) {
	childrenCount := treeData[treeDataIndex]
	metadataCount := treeData[treeDataIndex+1]
	treeDataIndex += 2

	// Add new node
	ID := t.AddNode(parentID)

	for i := 0; i < childrenCount; i++ { // For each child, call recursive
		treeDataIndex = instructionReadRec(t, ID, treeData, treeDataIndex)
	}
	t.AddMetadata(ID, treeData[treeDataIndex:treeDataIndex+metadataCount]) // Add metadata to node
	treeDataIndex += metadataCount

	return treeDataIndex
}

func main() {
	data, _ := ioutil.ReadFile("./input")
	treeData := []int{}
	for _, n := range s.Split(s.Trim(string(data), "\n"), " ") {
		nint, _ := sc.Atoi(n)
		treeData = append(treeData, nint)
	}

	sum, val := readTree(treeData)
	output := "sum of treeData: " + sc.Itoa(sum) + "\nvalue of root node: " + sc.Itoa(val) + "\n"
	fmt.Println(output)
	ioutil.WriteFile("./output", []byte(output), 0644)
}
