package main

import (
	"fmt"
	"math"
	"strings"
)

type Day11 struct{}

type Day11Galaxy struct {
	x    int
	y    int
	name string
}

func (d Day11) Part1(input string) {
	lines := strings.Split(input, "\r\n")

	var verticallyExpanded []string
	for _, line := range lines {
		containsGalaxy := false
		for _, character := range line {
			if character == '#' {
				containsGalaxy = true
				break
			}
		}
		if containsGalaxy {
			verticallyExpanded = append(verticallyExpanded, line)
		} else {
			verticallyExpanded = append(verticallyExpanded, line)
			verticallyExpanded = append(verticallyExpanded, line)
		}
	}
	lines = verticallyExpanded

	for i := 0; i < len(lines[0]); i++ {
		containsGalaxy := false
		for _, line := range lines {
			if line[i] == '#' {
				containsGalaxy = true
				break
			}
		}
		if !containsGalaxy {
			for j, line := range lines {
				lines[j] = line[:i] + ".." + line[i+1:]
			}
			i += 1
		}
	}

	var galaxies []Day11Galaxy
	for y, line := range lines {
		for x, character := range line {
			if character == '#' {
				galaxies = append(galaxies, Day11Galaxy{
					x:    x,
					y:    y,
					name: fmt.Sprintf("%d,%d", x, y),
				})
			}
		}
	}

	var sum int
	for _, galaxy1 := range galaxies {
		for _, galaxy2 := range galaxies {
			if galaxy1.name == galaxy2.name {
				continue
			}

			distanceY := math.Abs(float64(galaxy2.y - galaxy1.y))
			distanceX := math.Abs(float64(galaxy2.x - galaxy1.x))
			distance := int(distanceY + distanceX)
			sum += distance
		}
	}

	fmt.Println(sum / 2)
}

func (d Day11) Part2(input string) {
	lines := strings.Split(input, "\r\n")
	howMuchLarger := float64(1000000)

	var verticallyExpanded []int
	for y, line := range lines {
		containsGalaxy := false
		for _, character := range line {
			if character == '#' {
				containsGalaxy = true
				break
			}
		}
		if !containsGalaxy {
			verticallyExpanded = append(verticallyExpanded, y)
		}
	}

	var horizontallyExpanded []int
	for x := 0; x < len(lines[0]); x++ {
		containsGalaxy := false
		for _, line := range lines {
			if line[x] == '#' {
				containsGalaxy = true
				break
			}
		}
		if !containsGalaxy {
			horizontallyExpanded = append(horizontallyExpanded, x)
		}
	}

	var galaxies []Day11Galaxy
	for y, line := range lines {
		for x, character := range line {
			if character == '#' {
				galaxies = append(galaxies, Day11Galaxy{
					x:    x,
					y:    y,
					name: fmt.Sprintf("%d,%d", x, y),
				})
			}
		}
	}

	var sum int
	for _, galaxy1 := range galaxies {
		for _, galaxy2 := range galaxies {
			if galaxy1.name == galaxy2.name {
				continue
			}

			var horizontalExpansions int
			for _, x := range horizontallyExpanded {
				if galaxy1.x > galaxy2.x && x < galaxy1.x && x > galaxy2.x ||
					galaxy1.x < galaxy2.x && x > galaxy1.x && x < galaxy2.x {
					horizontalExpansions++
				}
			}

			var verticalExpansions int
			for _, y := range verticallyExpanded {
				if galaxy1.y > galaxy2.y && y < galaxy1.y && y > galaxy2.y ||
					galaxy1.y < galaxy2.y && y > galaxy1.y && y < galaxy2.y {
					verticalExpansions++
				}
			}

			distanceY := math.Abs(float64(galaxy2.y-galaxy1.y)) - float64(verticalExpansions) + howMuchLarger*float64(verticalExpansions)
			distanceX := math.Abs(float64(galaxy2.x-galaxy1.x)) - float64(horizontalExpansions) + howMuchLarger*float64(horizontalExpansions)
			distance := int(distanceY + distanceX)
			sum += distance
		}
	}

	fmt.Println(sum / 2)
}
