package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Day2 struct {
}

func (d *Day2) Part1(input string) {
	maxRedCubes := 12
	maxGreenCubes := 13
	maxBlueCubes := 14

	lines := strings.Split(input, "\r\n")

	var sum int
	for _, line := range lines {
		valid := true

		parts := strings.Split(line, ":")
		number, err := strconv.Atoi(strings.TrimLeft(parts[0], "Game "))
		if err != nil {
			panic(err)
		}

		rounds := strings.Split(parts[1], ";")
		for _, round := range rounds {
			pulls := strings.Split(round, ",")
			for _, pull := range pulls {
				parts = strings.Split(strings.TrimSpace(pull), " ")
				count, err := strconv.Atoi(parts[0])
				if err != nil {
					panic(err)
				}
				color := strings.TrimSpace(parts[1])
				if color == "red" && count > maxRedCubes {
					valid = false
				}
				if color == "green" && count > maxGreenCubes {
					valid = false
				}
				if color == "blue" && count > maxBlueCubes {
					valid = false
				}
			}
		}

		if valid {
			sum += number
		}
	}

	fmt.Println(sum)
}

func (d *Day2) Part2(input string) {
	lines := strings.Split(input, "\r\n")

	var sum int
	for _, line := range lines {
		var minRedCubes int
		var minGreenCubes int
		var minBlueCubes int

		parts := strings.Split(line, ":")

		rounds := strings.Split(parts[1], ";")
		for _, round := range rounds {
			pulls := strings.Split(round, ",")
			for _, pull := range pulls {
				parts = strings.Split(strings.TrimSpace(pull), " ")
				count, err := strconv.Atoi(parts[0])
				if err != nil {
					panic(err)
				}
				color := strings.TrimSpace(parts[1])
				if color == "red" && count > minRedCubes {
					minRedCubes = count
				}
				if color == "green" && count > minGreenCubes {
					minGreenCubes = count
				}
				if color == "blue" && count > minBlueCubes {
					minBlueCubes = count
				}
			}
		}

		sum += minRedCubes * minGreenCubes * minBlueCubes
	}

	fmt.Println(sum)
}
