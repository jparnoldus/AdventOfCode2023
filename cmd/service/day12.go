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

	var amounts []int
	for i, line := range lines {
		fmt.Print(i, ": ")

		var sum int
		parts := strings.Split(line, " ")
		filter := fmt.Sprintf("%s?%s?%s?%s?%s", parts[0], parts[0], parts[0], parts[0], parts[0])
		linePattern := fmt.Sprintf("%s,%s,%s,%s,%s", parts[1], parts[1], parts[1], parts[1], parts[1])
		//filter := parts[0]
		//linePattern := parts[1]
		numberStrings := strings.Split(linePattern, ",")
		var numbers []int
		for _, numberString := range numberStrings {
			number, _ := strconv.Atoi(numberString)
			numbers = append(numbers, number)
		}

		var amountOfFixed int
		var wildcardLocations []int
		for k, character := range filter {
			if character == '?' {
				wildcardLocations = append(wildcardLocations, k)
			} else if character == '#' {
				amountOfFixed++
			}
		}

		var amountMissing int
		for _, number := range numbers {
			amountMissing += number
		}
		amountMissing -= amountOfFixed

		lastPatternAmount := 0
		patternIndex := -1
		addition := filter[:wildcardLocations[0]]
		var guess string
		var temp string
		for n, char := range addition {
			if n == 0 {
				temp = string(char)
			}
			if temp == "#" && string(char) == "#" {
				lastPatternAmount++
			} else if temp == "." && string(char) == "#" {
				lastPatternAmount = 1
				patternIndex++
			}
			temp = string(char)
			guess += string(char)
		}

		sum += d.Do(guess, "#", filter, numbers, lastPatternAmount, patternIndex, wildcardLocations, amountMissing)
		sum += d.Do(guess, ".", filter, numbers, lastPatternAmount, patternIndex, wildcardLocations, amountMissing)

		amounts = append(amounts, sum)
		fmt.Println(sum)
	}

	var sum int
	for _, amount := range amounts {
		sum += amount
	}

	fmt.Println(sum)
}

func (d Day12) Do(guess string, new string, filter string, linePattern []int, lastPatternAmount int, patternIndex int, wildcardLocations []int, missing int) int {
	if len(guess) > 0 {
		last := guess[len(guess)-1]
		guess += new

		if last == '#' && new == "#" {
			missing--
			lastPatternAmount++
			if lastPatternAmount > linePattern[patternIndex] {
				return 0
			}
		} else if last == '.' && new == "#" {
			missing--
			lastPatternAmount = 1
			patternIndex++
			if patternIndex > len(linePattern)-1 {
				return 0
			}
		} else if last == '#' && new == "." {
			if patternIndex > len(linePattern)-1 {
				return 0
			}
			if lastPatternAmount != linePattern[patternIndex] {
				return 0
			}
		}
	} else {
		guess += new
		if new == "#" {
			missing--
			lastPatternAmount = 1
			patternIndex++
		}
	}

	for _, location := range wildcardLocations {
		if location >= len(guess) {
			if location-len(guess) >= 1 {
				addition := filter[len(guess):location]
				for _, char := range addition {
					if new == "#" && string(char) == "#" {
						lastPatternAmount++
						if lastPatternAmount > linePattern[patternIndex] {
							return 0
						}
					} else if new == "." && string(char) == "#" {
						lastPatternAmount = 1
						patternIndex++
						if patternIndex > len(linePattern)-1 {
							return 0
						}
					} else if new == "#" && string(char) == "." {
						if patternIndex > len(linePattern)-1 {
							return 0
						}
						if lastPatternAmount != linePattern[patternIndex] {
							return 0
						}
					}
					new = string(char)
					guess += string(char)
				}
			}
			break
		}
	}

	if missing == 0 {
		guess += filter[len(guess):]
		guess = strings.ReplaceAll(guess, "?", ".")
		if len(guess) != len(filter) {
			return 0
		}

		return fullCheck(guess, filter, linePattern)
	}

	if len(guess) > wildcardLocations[len(wildcardLocations)-1] {
		guess += filter[wildcardLocations[len(wildcardLocations)-1]+1:]
		return fullCheck(guess, filter, linePattern)
	}

	if len(guess) == len(filter) || len(guess) > wildcardLocations[len(wildcardLocations)-1] {
		return fullCheck(guess, filter, linePattern)
	}

	var sum int
	if new == "#" && lastPatternAmount == linePattern[patternIndex] {
		sum += d.Do(guess, ".", filter, linePattern, lastPatternAmount, patternIndex, wildcardLocations, missing)
	} else if new == "#" && lastPatternAmount < linePattern[patternIndex] {
		sum += d.Do(guess, "#", filter, linePattern, lastPatternAmount, patternIndex, wildcardLocations, missing)
	} else if filter[len(guess)] == '?' {
		sum += d.Do(guess, "#", filter, linePattern, lastPatternAmount, patternIndex, wildcardLocations, missing)
		sum += d.Do(guess, ".", filter, linePattern, lastPatternAmount, patternIndex, wildcardLocations, missing)
	} else {
		sum += d.Do(guess, string(filter[len(guess)]), filter, linePattern, lastPatternAmount, patternIndex, wildcardLocations, missing)
	}
	return sum
}

func fullCheck(guess string, filter string, linePattern []int) int {
	for i, char1 := range guess {
		char2 := filter[i]
		if char2 == '?' {
			continue
		}
		if char1 != rune(char2) {
			return 0
		}
	}

	var test []int
	var current int
	for _, character := range guess {
		if character == '#' {
			current++
		} else {
			if current > 0 {
				test = append(test, current)
			}
			current = 0
		}
	}
	if current > 0 {
		test = append(test, current)
	}
	if len(test) != len(linePattern) {
		return 0
	}
	for i, number := range test {
		if number != linePattern[i] {
			return 0
		}
	}

	return 1
}
