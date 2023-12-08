package main

import (
	"fmt"
	"strings"
)

type Day8 struct{}

func (d Day8) Part1(input string) {
	lines := strings.Split(input, "\r\n")

	var directions []string
	network := make(map[string]map[string]string)
	for i, line := range lines {
		if line == "" {
			continue
		}
		if i == 0 {
			for _, character := range line {
				directions = append(directions, string(character))
			}
			continue
		}

		line = strings.ReplaceAll(line, " ", "")
		line = strings.ReplaceAll(line, "(", "")
		line = strings.ReplaceAll(line, ")", "")
		parts := strings.Split(line, "=")
		from := parts[0]
		to := strings.Split(parts[1], ",")
		network[from] = make(map[string]string)
		network[from]["L"] = to[0]
		network[from]["R"] = to[1]
	}

	var steps int
	var index int
	var current = "AAA"
	for current != "ZZZ" {
		steps++
		if directions[index] == "L" {
			current = network[current]["L"]
		} else {
			current = network[current]["R"]
		}
		index++
		if index == len(directions) {
			index = 0
		}
	}

	fmt.Println(steps)
}

func (d Day8) Part2(input string) {
	lines := strings.Split(input, "\r\n")

	var startingPoints []string
	var directions []string
	network := make(map[string]map[string]string)
	for i, line := range lines {
		if line == "" {
			continue
		}
		if i == 0 {
			for _, character := range line {
				directions = append(directions, string(character))
			}
			continue
		}

		line = strings.ReplaceAll(line, " ", "")
		line = strings.ReplaceAll(line, "(", "")
		line = strings.ReplaceAll(line, ")", "")
		parts := strings.Split(line, "=")
		from := parts[0]
		to := strings.Split(parts[1], ",")
		network[from] = make(map[string]string)
		network[from]["L"] = to[0]
		network[from]["R"] = to[1]

		if strings.HasSuffix(from, "A") {
			startingPoints = append(startingPoints, from)
		}
	}

	initial := len(startingPoints)
	var allSteps []int
	var steps int
	var index int
	for len(allSteps) != initial {
		steps++
		var newPoints []string
		for _, current := range startingPoints {
			if directions[index] == "L" {
				current = network[current]["L"]
			} else {
				current = network[current]["R"]
			}

			if strings.HasSuffix(current, "Z") {
				allSteps = append(allSteps, steps)
				continue
			}

			newPoints = append(newPoints, current)
		}
		startingPoints = newPoints
		index++
		if index == len(directions) {
			index = 0
		}
	}

	lcm := d.lcm(allSteps[0], allSteps[1])
	for i := 2; i < len(allSteps); i++ {
		lcm = d.lcm(lcm, allSteps[i])
	}

	fmt.Println(lcm)
}

func (d Day8) lcm(a, b int) int {
	return a * b / d.gcd(a, b)
}

func (d Day8) gcd(a, b int) int {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}
	return a
}
