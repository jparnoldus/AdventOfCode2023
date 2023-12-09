package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Day6 struct {
}

func (d Day6) Part1(input string) {
	lines := strings.Split(input, "\r\n")
	timesLine := lines[0]
	distancesLine := lines[1]

	timesUnfiltered := strings.Split(strings.TrimSpace(strings.Split(timesLine, ":")[1]), " ")
	var times []int
	for _, timeString := range timesUnfiltered {
		if timeString != "" {
			time, err := strconv.Atoi(timeString)
			if err != nil {
				panic(err)
			}
			times = append(times, time)
		}
	}

	distancesUnfiltered := strings.Split(strings.TrimSpace(strings.Split(distancesLine, ":")[1]), " ")
	var distances []int
	for _, distanceString := range distancesUnfiltered {
		if distanceString != "" {
			distance, err := strconv.Atoi(distanceString)
			if err != nil {
				panic(err)
			}
			distances = append(distances, distance)
		}
	}

	multiplied := 1
	for i, time := range times {
		distance := distances[i]
		var possibilities []int
		for i := 0; i < time; i++ {
			if i*(time-i) > distance {
				possibilities = append(possibilities, i)
			}
		}

		difference := possibilities[len(possibilities)-1] - possibilities[0] + 1
		multiplied *= difference
	}

	fmt.Println(multiplied)
}

func (d Day6) Part2(input string) {
	lines := strings.Split(input, "\r\n")
	timesLine := lines[0]
	distancesLine := lines[1]

	timeString := strings.ReplaceAll(strings.Split(timesLine, ":")[1], " ", "")
	time, _ := strconv.Atoi(timeString)

	distanceString := strings.ReplaceAll(strings.Split(distancesLine, ":")[1], " ", "")
	distance, _ := strconv.Atoi(distanceString)

	var possibilities []int
	for i := 0; i < time; i++ {
		if i*(time-i) > distance {
			possibilities = append(possibilities, i)
		}
	}
	difference := possibilities[len(possibilities)-1] - possibilities[0] + 1

	fmt.Println(difference)
}
