package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Day12 struct {
}

func (d Day12) Part1(input string) {
	lines := strings.Split(input, "\r\n")

	var sum int
	for _, line := range lines {
		parts := strings.Split(line, " ")
		filter := parts[0]
		linePattern := parts[1]

		maxLength := len(filter)
		for i := 0; i < int(math.Pow(float64(2), float64(maxLength))); i++ {
			binary := fmt.Sprintf("%0"+strconv.Itoa(maxLength)+"b", i)
			binary = strings.ReplaceAll(binary, "1", "#")
			binary = strings.ReplaceAll(binary, "0", ".")

			fits := true
			for j, char1 := range binary {
				char2 := string(filter[j])
				if char2 == "?" {
					continue
				}
				if string(char1) != char2 {
					fits = false
					break
				}
			}
			if fits {
				var pattern string
				var current int
				for _, character := range binary {
					if character == '#' {
						current++
					} else {
						if current > 0 {
							pattern += fmt.Sprintf(",%d", current)
						}
						current = 0
					}
				}
				if current > 0 {
					pattern += fmt.Sprintf(",%d", current)
				}
				pattern = strings.TrimPrefix(pattern, ",")

				if pattern == linePattern {
					sum++
				}
			}
		}
	}

	fmt.Println(sum)
}

func (d Day12) Part2(input string) {
	lines := strings.Split(input, "\r\n")

	var sum int
	for _, line := range lines {
		parts := strings.Split(line, " ")
		filter := fmt.Sprintf("%s?%s?%s?%s?%s", parts[0], parts[0], parts[0], parts[0], parts[0])
		linePattern := fmt.Sprintf("%s,%s,%s,%s,%s", parts[1], parts[1], parts[1], parts[1], parts[1])

		var amountOfFixed float64
		var amountOfWildcards float64
		for _, character := range filter {
			if character == '?' {
				amountOfWildcards++
			} else if character == '#' {
				amountOfFixed++
			}
		}

		var amountMissing float64
		numbers := strings.Split(linePattern, ",")
		for _, number := range numbers {
			temp, _ := strconv.Atoi(number)
			amountMissing += float64(temp)
		}
		amountMissing -= amountOfFixed

		regex := "%0" + strconv.Itoa(int(amountOfWildcards)) + "b"
		for i := 0.0; i < math.Pow(float64(2), amountOfWildcards); i++ {
			temp := filter
			binary := fmt.Sprintf(regex, i)
			for _, character := range binary {
				temp = strings.Replace(temp, "?", string(character), 1)
			}
		}

		var pattern string
		var current int
		for _, character := range filter {
			if character == '#' {
				current++
			} else {
				if current > 0 {
					pattern += fmt.Sprintf(",%d", current)
				}
				current = 0
			}
		}
		if current > 0 {
			pattern += fmt.Sprintf(",%d", current)
		}
		pattern = strings.TrimPrefix(pattern, ",")

		if pattern == linePattern {
			sum++
		}
	}

	fmt.Println("ok")
}
