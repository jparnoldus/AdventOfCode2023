package main

import (
	"fmt"
	"strings"
)

type Day10 struct {
}

type Day10Point struct {
	x         int
	y         int
	value     string
	direction string
}

func (d Day10) Part1(input string) {
	lines := strings.Split(input, "\r\n")

	var startingPoint Day10Point
	for y, line := range lines {
		for x, character := range line {
			if character == 'S' {
				startingPoint = Day10Point{x: x, y: y, value: string(character)}
			}
		}
	}

	var route []Day10Point
	route = append(route, startingPoint)

	for route[len(route)-1].value != "S" || len(route) == 1 {
		current := route[len(route)-1]
		switch route[len(route)-1].value {
		case "S":
			if string(lines[current.y+1][current.x]) == "|" || string(lines[current.y+1][current.x]) == "J" || string(lines[current.y+1][current.x]) == "L" {
				route = append(route, Day10Point{
					x:     current.x,
					y:     current.y + 1,
					value: string(lines[current.y+1][current.x]),
				})
			} else if string(lines[current.y-1][current.x]) == "|" || string(lines[current.y-1][current.x]) == "F" || string(lines[current.y-1][current.x]) == "7" {
				route = append(route, Day10Point{
					x:     current.x,
					y:     current.y - 1,
					value: string(lines[current.y-1][current.x]),
				})
			} else if string(lines[current.y][current.x+1]) == "-" || string(lines[current.y-1][current.x]) == "J" || string(lines[current.y-1][current.x]) == "7" {
				route = append(route, Day10Point{
					x:     current.x + 1,
					y:     current.y,
					value: string(lines[current.y][current.x+1]),
				})
			} else if string(lines[current.y][current.x-1]) == "-" || string(lines[current.y-1][current.x]) == "F" || string(lines[current.y-1][current.x]) == "L" {
				route = append(route, Day10Point{
					x:     current.x - 1,
					y:     current.y,
					value: string(lines[current.y][current.x-1]),
				})
			}
			continue
		case "-":
			previous := route[len(route)-2]
			if previous.x == current.x+1 {
				route = append(route, Day10Point{
					x:     current.x - 1,
					y:     current.y,
					value: string(lines[current.y][current.x-1]),
				})
			} else {
				route = append(route, Day10Point{
					x:     current.x + 1,
					y:     current.y,
					value: string(lines[current.y][current.x+1]),
				})
			}
			continue
		case "|":
			previous := route[len(route)-2]
			if previous.y == current.y+1 {
				route = append(route, Day10Point{
					x:     current.x,
					y:     current.y - 1,
					value: string(lines[current.y-1][current.x]),
				})
			} else {
				route = append(route, Day10Point{
					x:     current.x,
					y:     current.y + 1,
					value: string(lines[current.y+1][current.x]),
				})
			}
			continue
		case "J":
			previous := route[len(route)-2]
			if previous.y == current.y-1 {
				route = append(route, Day10Point{
					x:     current.x - 1,
					y:     current.y,
					value: string(lines[current.y][current.x-1]),
				})
			} else {
				route = append(route, Day10Point{
					x:     current.x,
					y:     current.y - 1,
					value: string(lines[current.y-1][current.x]),
				})
			}
			continue
		case "L":
			previous := route[len(route)-2]
			if previous.y == current.y-1 {
				route = append(route, Day10Point{
					x:     current.x + 1,
					y:     current.y,
					value: string(lines[current.y][current.x+1]),
				})
			} else {
				route = append(route, Day10Point{
					x:     current.x,
					y:     current.y - 1,
					value: string(lines[current.y-1][current.x]),
				})
			}
			continue
		case "7":
			previous := route[len(route)-2]
			if previous.y == current.y+1 {
				route = append(route, Day10Point{
					x:     current.x - 1,
					y:     current.y,
					value: string(lines[current.y][current.x-1]),
				})
			} else {
				route = append(route, Day10Point{
					x:     current.x,
					y:     current.y + 1,
					value: string(lines[current.y+1][current.x]),
				})
			}
			continue
		case "F":
			previous := route[len(route)-2]
			if previous.y == current.y+1 {
				route = append(route, Day10Point{
					x:     current.x + 1,
					y:     current.y,
					value: string(lines[current.y][current.x+1]),
				})
			} else {
				route = append(route, Day10Point{
					x:     current.x,
					y:     current.y + 1,
					value: string(lines[current.y+1][current.x]),
				})
			}
			continue
		}
	}

	farthest := (len(route) - 1) / 2

	fmt.Println(farthest)
}

func (d Day10) Part2(input string) {
	lines := strings.Split(input, "\r\n")

	var startingPoint Day10Point
	for y, line := range lines {
		for x, character := range line {
			if character == 'S' {
				startingPoint = Day10Point{x: x, y: y, value: string(character)}
			}
		}
	}

	var route []Day10Point
	route = append(route, startingPoint)

	for route[len(route)-1].value != "S" || len(route) == 1 {
		current := route[len(route)-1]
		switch route[len(route)-1].value {
		case "S":
			if string(lines[current.y+1][current.x]) == "|" || string(lines[current.y+1][current.x]) == "J" || string(lines[current.y+1][current.x]) == "L" {
				route = append(route, Day10Point{
					x:     current.x,
					y:     current.y + 1,
					value: string(lines[current.y+1][current.x]),
				})
			} else if string(lines[current.y-1][current.x]) == "|" || string(lines[current.y-1][current.x]) == "F" || string(lines[current.y-1][current.x]) == "7" {
				route = append(route, Day10Point{
					x:     current.x,
					y:     current.y - 1,
					value: string(lines[current.y-1][current.x]),
				})
			} else if string(lines[current.y][current.x+1]) == "-" || string(lines[current.y-1][current.x]) == "J" || string(lines[current.y-1][current.x]) == "7" {
				route = append(route, Day10Point{
					x:     current.x + 1,
					y:     current.y,
					value: string(lines[current.y][current.x+1]),
				})
			} else if string(lines[current.y][current.x-1]) == "-" || string(lines[current.y-1][current.x]) == "F" || string(lines[current.y-1][current.x]) == "L" {
				route = append(route, Day10Point{
					x:     current.x - 1,
					y:     current.y,
					value: string(lines[current.y][current.x-1]),
				})
			}
			continue
		case "-":
			previous := route[len(route)-2]
			if previous.x == current.x+1 {
				route = append(route, Day10Point{
					x:     current.x - 1,
					y:     current.y,
					value: string(lines[current.y][current.x-1]),
				})
			} else {
				route = append(route, Day10Point{
					x:     current.x + 1,
					y:     current.y,
					value: string(lines[current.y][current.x+1]),
				})
			}
			continue
		case "|":
			previous := route[len(route)-2]
			if previous.y == current.y+1 {
				route = append(route, Day10Point{
					x:     current.x,
					y:     current.y - 1,
					value: string(lines[current.y-1][current.x]),
				})
			} else {
				route = append(route, Day10Point{
					x:     current.x,
					y:     current.y + 1,
					value: string(lines[current.y+1][current.x]),
				})
			}
			continue
		case "J":
			previous := route[len(route)-2]
			if previous.y == current.y-1 {
				route = append(route, Day10Point{
					x:     current.x - 1,
					y:     current.y,
					value: string(lines[current.y][current.x-1]),
				})
			} else {
				route = append(route, Day10Point{
					x:     current.x,
					y:     current.y - 1,
					value: string(lines[current.y-1][current.x]),
				})
			}
			continue
		case "L":
			previous := route[len(route)-2]
			if previous.y == current.y-1 {
				route = append(route, Day10Point{
					x:     current.x + 1,
					y:     current.y,
					value: string(lines[current.y][current.x+1]),
				})
			} else {
				route = append(route, Day10Point{
					x:     current.x,
					y:     current.y - 1,
					value: string(lines[current.y-1][current.x]),
				})
			}
			continue
		case "7":
			previous := route[len(route)-2]
			if previous.y == current.y+1 {
				route = append(route, Day10Point{
					x:     current.x - 1,
					y:     current.y,
					value: string(lines[current.y][current.x-1]),
				})
			} else {
				route = append(route, Day10Point{
					x:     current.x,
					y:     current.y + 1,
					value: string(lines[current.y+1][current.x]),
				})
			}
			continue
		case "F":
			previous := route[len(route)-2]
			if previous.y == current.y+1 {
				route = append(route, Day10Point{
					x:     current.x + 1,
					y:     current.y,
					value: string(lines[current.y][current.x+1]),
				})
			} else {
				route = append(route, Day10Point{
					x:     current.x,
					y:     current.y + 1,
					value: string(lines[current.y+1][current.x]),
				})
			}
			continue
		}
	}

	for i, point := range route {
		if i == 0 {
			continue
		}
		if point.value == "S" {
			break
		}

		if i == 1 {
			switch point.value {
			case "-":
				point.direction = "T"
				break
			case "|":
				point.direction = "R"
				break
			case "J":
				point.direction = "O"
				break
			case "L":
				point.direction = "O"
				break
			case "7":
				point.direction = "O"
				break
			case "F":
				point.direction = "O"
				break
			}
			route[i] = point
			continue
		}

		previous := route[i-1]
		switch previous.value {
		case "-":
			switch point.value {
			case "-":
				point.direction = previous.direction
				break
			case "J", "L":
				if previous.direction == "T" {
					point.direction = "I"
				} else {
					point.direction = "O"
				}
				break
			case "7", "F":
				if previous.direction == "T" {
					point.direction = "O"
				} else {
					point.direction = "I"
				}
				break
			}
			break
		case "|":
			switch point.value {
			case "|":
				point.direction = previous.direction
				break
			case "J", "7":
				if previous.direction == "R" {
					point.direction = "O"
				} else {
					point.direction = "I"
				}
				break
			case "L", "F":
				if previous.direction == "R" {
					point.direction = "I"
				} else {
					point.direction = "O"
				}
				break
			}
			break
		case "J":
			switch point.value {
			case "-":
				if previous.direction == "O" {
					point.direction = "B"
				} else {
					point.direction = "T"
				}
				break
			case "|":
				if previous.direction == "O" {
					point.direction = "R"
				} else {
					point.direction = "L"
				}
				break
			case "L", "7":
				point.direction = previous.direction
				break
			case "F":
				if previous.direction == "O" {
					point.direction = "I"
				} else {
					point.direction = "O"
				}
				break
			}
			break
		case "L":
			switch point.value {
			case "-":
				if previous.direction == "O" {
					point.direction = "B"
				} else {
					point.direction = "T"
				}
				break
			case "|":
				if previous.direction == "O" {
					point.direction = "L"
				} else {
					point.direction = "R"
				}
				break
			case "J", "F":
				point.direction = previous.direction
				break
			case "7":
				if previous.direction == "O" {
					point.direction = "I"
				} else {
					point.direction = "O"
				}
				break
			}
			break
		case "7":
			switch point.value {
			case "-":
				if previous.direction == "O" {
					point.direction = "T"
				} else {
					point.direction = "B"
				}
				break
			case "|":
				if previous.direction == "O" {
					point.direction = "R"
				} else {
					point.direction = "L"
				}
				break
			case "J", "F":
				point.direction = previous.direction
				break
			case "L":
				if previous.direction == "O" {
					point.direction = "I"
				} else {
					point.direction = "O"
				}
				break
			}
			break
		case "F":
			switch point.value {
			case "-":
				if previous.direction == "O" {
					point.direction = "T"
				} else {
					point.direction = "B"
				}
				break
			case "|":
				if previous.direction == "O" {
					point.direction = "L"
				} else {
					point.direction = "R"
				}
				break
			case "J":
				if previous.direction == "O" {
					point.direction = "I"
				} else {
					point.direction = "O"
				}
				break
			case "L", "7":
				point.direction = previous.direction
				break
			}
			break
		}
		route[i] = point
	}

	for _, point := range route {
		if point.value == "S" {
			continue
		}

		switch point.direction {
		case "T":
			if !d.IsPartOfRoute(point.x, point.y-1, route) && d.IsInBounds(point.x, point.y-1, lines) {
				lines[point.y-1] = lines[point.y-1][:point.x] + "O" + lines[point.y-1][point.x+1:]
			}
			break
		case "B":
			if !d.IsPartOfRoute(point.x, point.y+1, route) && d.IsInBounds(point.x, point.y+1, lines) {
				lines[point.y+1] = lines[point.y+1][:point.x] + "O" + lines[point.y+1][point.x+1:]
			}
			break
		case "L":
			if !d.IsPartOfRoute(point.x-1, point.y, route) && d.IsInBounds(point.x-1, point.y, lines) {
				lines[point.y] = lines[point.y][:point.x-1] + "O" + lines[point.y][point.x:]
			}
			break
		case "R":
			if !d.IsPartOfRoute(point.x+1, point.y, route) && d.IsInBounds(point.x+1, point.y, lines) {
				lines[point.y] = lines[point.y][:point.x+1] + "O" + lines[point.y][point.x+2:]
			}
			break
		case "O":
			switch point.value {
			case "J":
				if !d.IsPartOfRoute(point.x, point.y+1, route) && d.IsInBounds(point.x, point.y+1, lines) {
					lines[point.y+1] = lines[point.y+1][:point.x] + "O" + lines[point.y+1][point.x+1:]
				}
				if !d.IsPartOfRoute(point.x+1, point.y, route) && d.IsInBounds(point.x+1, point.y, lines) {
					lines[point.y] = lines[point.y][:point.x+1] + "O" + lines[point.y][point.x+2:]
				}
				if !d.IsPartOfRoute(point.x+1, point.y+1, route) && d.IsInBounds(point.x+1, point.y+1, lines) {
					lines[point.y+1] = lines[point.y+1][:point.x+1] + "O" + lines[point.y+1][point.x+2:]
				}
				break
			case "L":
				if !d.IsPartOfRoute(point.x-1, point.y, route) && d.IsInBounds(point.x-1, point.y, lines) {
					lines[point.y] = lines[point.y][:point.x-1] + "O" + lines[point.y][point.x:]
				}
				if !d.IsPartOfRoute(point.x-1, point.y+1, route) && d.IsInBounds(point.x-1, point.y+1, lines) {
					lines[point.y+1] = lines[point.y+1][:point.x-1] + "O" + lines[point.y+1][point.x:]
				}
				if !d.IsPartOfRoute(point.x, point.y+1, route) && d.IsInBounds(point.x, point.y+1, lines) {
					lines[point.y+1] = lines[point.y+1][:point.x] + "O" + lines[point.y+1][point.x+1:]
				}
				break
			case "7":
				if !d.IsPartOfRoute(point.x, point.y-1, route) && d.IsInBounds(point.x, point.y-1, lines) {
					lines[point.y-1] = lines[point.y-1][:point.x] + "O" + lines[point.y-1][point.x+1:]
				}
				if !d.IsPartOfRoute(point.x+1, point.y, route) && d.IsInBounds(point.x+1, point.y, lines) {
					lines[point.y] = lines[point.y][:point.x+1] + "O" + lines[point.y][point.x+2:]
				}
				if !d.IsPartOfRoute(point.x+1, point.y-1, route) && d.IsInBounds(point.x+1, point.y-1, lines) {
					lines[point.y-1] = lines[point.y-1][:point.x+1] + "O" + lines[point.y-1][point.x+2:]
				}
				break
			case "F":
				if !d.IsPartOfRoute(point.x-1, point.y, route) && d.IsInBounds(point.x-1, point.y, lines) {
					lines[point.y] = lines[point.y][:point.x-1] + "O" + lines[point.y][point.x:]
				}
				if !d.IsPartOfRoute(point.x-1, point.y-1, route) && d.IsInBounds(point.x-1, point.y-1, lines) {
					lines[point.y-1] = lines[point.y-1][:point.x-1] + "O" + lines[point.y-1][point.x:]
				}
				if !d.IsPartOfRoute(point.x, point.y-1, route) && d.IsInBounds(point.x, point.y-1, lines) {
					lines[point.y-1] = lines[point.y-1][:point.x] + "O" + lines[point.y-1][point.x+1:]
				}
				break
			}
			break
		}
	}

	changed := 1
	for changed != 0 {
		changed = 0
		for y, line := range lines {
			for x, character := range line {
				if character != 'O' {
					continue
				}
				if y+1 != len(lines) && lines[y+1][x] != 'O' && !d.IsPartOfRoute(x, y+1, route) {
					lines[y+1] = lines[y+1][:x] + "O" + lines[y+1][x+1:]
					changed++
				}
				if y-1 != -1 && lines[y-1][x] != 'O' && !d.IsPartOfRoute(x, y-1, route) {
					lines[y-1] = lines[y-1][:x] + "O" + lines[y-1][x+1:]
					changed++
				}
				if x+1 != len(line) && lines[y][x+1] != 'O' && !d.IsPartOfRoute(x+1, y, route) {
					lines[y] = lines[y][:x+1] + "O" + lines[y][x+2:]
					changed++
				}
				if x-1 != -1 && lines[y][x-1] != 'O' && !d.IsPartOfRoute(x-1, y, route) {
					lines[y] = lines[y][:x-1] + "O" + lines[y][x:]
					changed++
				}
			}
		}
	}

	var count int
	for y, line := range lines {
		for x, character := range line {
			if !d.IsPartOfRoute(x, y, route) && string(character) != "O" {
				count++
			}
		}
	}

	fmt.Println(count)
}

func (d Day10) IsInBounds(x, y int, lines []string) bool {
	return y >= 0 && y < len(lines) && x >= 0 && x < len(lines[y])
}

func (d Day10) IsPartOfRoute(x, y int, route []Day10Point) bool {
	isPartOfRoute := false
	for _, point := range route {
		if point.x == x && point.y == y {
			isPartOfRoute = true
			break
		}
	}

	return isPartOfRoute
}
