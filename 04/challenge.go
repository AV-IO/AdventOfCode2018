package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	sc "strconv"
	s "strings"
)

// funcName :
//  	parameters:
//  		rawLogs: list of unparsed logs with timestamps and what action was performed
//  	return values:
//  		retA: description
func funcName(rawLogs []string) (retA int) {
	sort.Strings(rawLogs)

	// parsing logs into usable structure
	var parsedLogs [][]string
	for i := range rawLogs {
		rawLogs[i] = s.Replace(rawLogs[i], "[", "", -1)
		rawLogs[i] = s.Replace(rawLogs[i], "]", "", -1)
		parsedLogs = append(parsedLogs, s.Split(rawLogs[i], " "))
	}

	// finding how much time each guard spent sleeping
	// Somewhat state-machine like
	// This could probably be better done by traversing lines up and down, but I want to do it this way
	// Maybe this would be solved better more functionally
	guardTimes := make(map[string][]int)
	currentGuard := ""
	time := 0
	for _, l := range parsedLogs {
		switch l[2] {
		case "Guard":
			currentGuard = l[3]
		case "falls":
			time, _ = sc.Atoi(l[1][3:])
		case "wakes":
			awakeTime, _ := sc.Atoi(l[1][3:])
			for ; time < awakeTime; time++ {
				guardTimes[currentGuard] = append(guardTimes[currentGuard], time)
			}
		}
	}

	// finding which guard spent the most time sleeping
	time = 0
	for key, val := range guardTimes {
		if len(val) > time {
			currentGuard = key
			time = len(val)
		}
	}

	// finding most common hour asleep
	sort.Ints(guardTimes[currentGuard])
	for i, _ := range guardTimes[currentGuard] {

		guardTimes[currentGuard][i]
		if count > time {
			time = count
		}
	}

	return
}

func main() {
	data, _ := ioutil.ReadFile("./input")
	retA := funcName(s.Split(string(data), "\n"))
	output := "retA: " + sc.Itoa(retA) + "\n"
	fmt.Println(output)
	ioutil.WriteFile("./output", []byte(output), 0644)
}
