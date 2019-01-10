package main

import (
	"../10/sky"
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	sc "strconv"
	str "strings"
	"time"
)

func findEndSky(s *sky.Sky) (endSky string) {
	fmt.Println(
		"Keyboard Input:\n" +
			"\tSpace: Play/Pause\n" +
			"\t    f: forward step\n" +
			"\t    r: reverse step\n" +
			"\t    d: step size down\n" +
			"\t    u: step size up\n" +
			"\t    e: exit\n" +
			"\nPress Space to start\n\n",
	)

	statusChan := make(chan string)
	done := make(chan bool)

	go func() {
		stepSize := 8
		playing := false
		for {
			select {
			case instruction, stayAlive := <-statusChan: //TODO: check if non-blocking close will work.
				if !stayAlive {
					done <- true
					return
				}
				switch instruction {
				case " ":
					if playing {
						playing = false
					} else {
						playing = true
					}
				case "f":
					if stepSize < 0 {
						stepSize *= -1
					}
					s.PassTime(stepSize)
				case "r":
					if stepSize > 0 {
						stepSize *= -1
					}
					s.PassTime(stepSize)
				case "d":
					stepSize /= 2
				case "u":
					stepSize *= 2
				}
			default:
				if playing {
					s.PassTime(stepSize)
					time.Sleep(100 * time.Millisecond) // helps run at a watchable timeframe
				}
			}
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		switch scanner.Text() {
		case " ", "f", "r", "d", "u":
			statusChan <- scanner.Text()
		case "e":
			close(statusChan)
			break
		}
	}

	<-done // wait for all changes to be made

	return s.ToString()
}

func readVectors(vectorlist []string) (endSky string) {
	s := new(sky.Sky)
	for _, v := range vectorlist {
		px, _ := sc.Atoi(str.Trim(v[str.Index(v, "<"):str.Index(v, ",")], " "))
		py, _ := sc.Atoi(str.Trim(v[str.Index(v, ","):str.Index(v, ">")], " "))
		vx, _ := sc.Atoi(str.Trim(v[str.LastIndex(v, "<"):str.LastIndex(v, ",")], " "))
		vy, _ := sc.Atoi(str.Trim(v[str.LastIndex(v, "<"):str.LastIndex(v, ",")], " "))
		s.AddVector(sky.Vector{PositionX: px, PositionY: py, VelocityX: vx, VelocityY: vy})
	}

	return findEndSky(s)
}

func main() {
	data, _ := ioutil.ReadFile("./input")
	endSky := readVectors(str.Split(string(data), "\n"))
	output := "" + endSky + "\n"
	fmt.Println(output)
	ioutil.WriteFile("./output", []byte(output), 0644)
}
