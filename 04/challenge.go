package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	sc "strconv"
	s "strings"
)

// findSleepTime :
//  	parameters:
//  		rawLogs: list of unparsed logs with timestamps and what action was performed
//  	return values:
//  		mostAsleep: description
func findSleepTime(rawLogs []string) (mostAsleep int, frequentMinute int) {
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
	// This could probably be better done by traversing lines up and down
	// Maybe this would be solved better more functionally
	guardTimes := make(map[string]map[int]int)
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
				if _, guardExists := guardTimes[currentGuard]; guardExists { // if guard exists
					guardTimes[currentGuard][time]++
				} else {
					guardTimes[currentGuard] = map[int]int{time: 1}
				}
			}
		}
	}

	// finding which guard spent the most time sleeping
	time = 0
	frequentCount := 0
	for guard, timeMap := range guardTimes {
		totalTime := 0
		for minute, timeCount := range timeMap {
			totalTime += timeCount         // pt. 1
			if timeCount > frequentCount { // pt. 2
				guardID, _ := sc.Atoi(guard[1:])
				frequentMinute = minute * guardID
				frequentCount = timeCount
			}
		}
		if totalTime > time {
			currentGuard = guard
			time = totalTime
		}
	}

	// finding most common hour asleep
	time = 0
	for minute, count := range guardTimes[currentGuard] {
		if count > time {
			mostAsleep = minute
			time = count
		}
	}

	//
	guardID, _ := sc.Atoi(currentGuard[1:])
	mostAsleep *= guardID

	return
}

func main() {
	data, _ := ioutil.ReadFile("./input")
	mostAsleep, frequentMinute := findSleepTime(s.Split(string(data), "\n"))
	output := "most time asleep: " + sc.Itoa(mostAsleep) + "\nmost frequently asleep minute: " + sc.Itoa(frequentMinute) + "\n"
	fmt.Println(output)
	ioutil.WriteFile("./output", []byte(output), 0644)
}
