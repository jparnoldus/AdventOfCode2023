package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Day3 struct {
}

type Point2D struct {
	x int
	y int
}

func (d *Day3) Part1(input string) {
	lines := strings.Split(input, "\r\n")

	numbers := "0123456789"

	numberLocations := make(map[Point2D]string)
	var schematic [][]string
	for y, line := range lines {
		characters := strings.Split(line, "")

		var current string
		for x, character := range characters {
			if strings.ContainsAny(character, numbers) {
				current += character
			} else if current != "" {
				point := Point2D{x - len(current), y}
				numberLocations[point] = current
				current = ""
			}
		}
		if current != "" {
			point := Point2D{len(characters) - len(current), y}
			numberLocations[point] = current
		}

		schematic = append(schematic, characters)
	}

	outboundX := len(schematic[0])
	outboundY := len(schematic)

	var sum int
	notSymbols := numbers + "."
	for point2D, numberString := range numberLocations {
		adjacent := false
		numberStringLength := len(numberString)

		for y := point2D.y - 1; y < point2D.y+2; y++ {
			for x := point2D.x - 1; x < point2D.x+numberStringLength+1; x++ {
				if x < 0 || y < 0 || x >= outboundX || y >= outboundY {
					continue
				}
				if !strings.ContainsAny(schematic[y][x], notSymbols) {
					adjacent = true
				}
			}
		}

		if adjacent {
			number, err := strconv.Atoi(numberString)
			if err != nil {
				panic(err)
			}
			sum += number
		}
	}

	fmt.Println(sum)
}

func (d *Day3) Part2(input string) {
	lines := strings.Split(input, "\r\n")

	numbers := "0123456789"

	numberLocations := make(map[Point2D]string)
	var gearLocations []Point2D
	var schematic [][]string
	for y, line := range lines {
		characters := strings.Split(line, "")

		var current string
		for x, character := range characters {
			if character == "*" {
				gearLocations = append(gearLocations, Point2D{x, y})
			}

			if strings.ContainsAny(character, numbers) {
				current += character
			} else if current != "" {
				point := Point2D{x - len(current), y}
				numberLocations[point] = current
				current = ""
			}
		}
		if current != "" {
			point := Point2D{len(characters) - len(current), y}
			numberLocations[point] = current
		}

		schematic = append(schematic, characters)
	}

	outboundX := len(schematic[0])
	outboundY := len(schematic)

	var sum int
	for _, point2D := range gearLocations {
		var adjacentNumbers []string

		for y := point2D.y - 1; y < point2D.y+2; y++ {
			for x := point2D.x - 3; x < point2D.x+2; x++ {
				if x < 0 || y < 0 || x >= outboundX || y >= outboundY {
					continue
				}

				point := Point2D{x, y}
				checked, exists := numberLocations[point]
				if exists && x+len(checked) >= point2D.x {
					adjacentNumbers = append(adjacentNumbers, checked)
				}
			}
		}

		var multiplied int
		if len(adjacentNumbers) > 1 {
			multiplied = 1
			for _, number := range adjacentNumbers {
				parsed, err := strconv.Atoi(number)
				if err != nil {
					panic(err)
				}
				multiplied *= parsed
			}
		}
		sum += multiplied
	}

	fmt.Println(sum)
}
