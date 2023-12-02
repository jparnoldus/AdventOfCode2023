package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Day1 struct {
}

func (d *Day1) Part1(input string) {
	lines := strings.Split(input, "\n")

	var numbersPerLines [][]string
	for _, line := range lines {
		var numbersPerLine []string
		for _, character := range line {
			if _, err := strconv.Atoi(string(character)); err == nil {
				numbersPerLine = append(numbersPerLine, string(character))
			}
		}
		numbersPerLines = append(numbersPerLines, numbersPerLine)
	}

	var sum int
	for _, line := range numbersPerLines {
		numberString := line[0] + line[len(line)-1]
		number, err := strconv.Atoi(numberString)
		if err != nil {
			panic(err)
		}
		sum += number
	}

	fmt.Println(sum)
}

func (d *Day1) Part2(input string) {
	lines := strings.Split(input, "\n")

	options := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	var firstPerLine []string
	for _, line := range lines {
		var first string
		for i, character := range line {
			if _, err := strconv.Atoi(string(character)); err == nil {
				first = string(character)
				break
			}

			tested := line[i:]
			for i, option := range options {
				if strings.HasPrefix(tested, option) {
					first = strconv.Itoa(i + 1)
					break
				}
			}
			if first != "" {
				break
			}
		}

		firstPerLine = append(firstPerLine, first)
	}

	var lastPerLine []string
	for _, line := range lines {
		length := len(line)
		var last string
		for i := range line {
			if _, err := strconv.Atoi(string(line[length-1-i])); err == nil {
				last = string(line[length-1-i])
				break
			}

			tested := line[:length-1-i]
			for i, option := range options {
				if strings.HasSuffix(tested, option) {
					last = strconv.Itoa(i + 1)
					break
				}
			}
			if last != "" {
				break
			}
		}

		lastPerLine = append(lastPerLine, last)
	}

	var sum int
	for i := range lines {
		number, err := strconv.Atoi(firstPerLine[i] + lastPerLine[i])
		if err != nil {
			panic(err)
		}
		sum += number
	}

	fmt.Println(sum)
}
