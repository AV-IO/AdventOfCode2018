package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// funcName :
//  	parameters:
//  		paramA: description
//  	return values:
//  		retA: description
func funcName(paramA []string) (retA int) {

	return
}

func main() {
	data, _ := ioutil.ReadFile("./input")
	retA := funcName(strings.Split(string(data), "\r\n"))
	output := "retA: " + strconv.Itoa(retA) + "\n"
	fmt.Println(output)
	ioutil.WriteFile("./output", []byte(output), 0644)
}
