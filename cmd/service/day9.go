package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Day9 struct {
}

func (d Day9) Part1(input string) {
	lines := strings.Split(input, "\r\n")

	var sum int
	for _, line := range lines {
		numberStrings := strings.Split(line, " ")

		var numbers []int
		for _, numberString := range numberStrings {
			number, err := strconv.Atoi(numberString)
			if err != nil {
				panic(err)
			}
			numbers = append(numbers, number)
		}

		var sequences [][]int
		sequences = append(sequences, numbers)
		for !d.isAllZeroes(sequences[len(sequences)-1]) {
			sequences = append(sequences, d.getDifferences(sequences[len(sequences)-1]))
		}

		sequences[len(sequences)-1] = append(sequences[len(sequences)-1], 0)

		for i := len(sequences) - 2; i > -1; i-- {
			currentSequence := sequences[i]
			previousSequence := sequences[i+1]
			newValue := currentSequence[len(previousSequence)-1] + previousSequence[len(previousSequence)-1]
			sequences[i] = append(sequences[i], newValue)
		}

		sum += sequences[0][len(sequences[0])-1]
	}

	fmt.Println(sum)
}

func (d Day9) getDifferences(values []int) []int {
	var differences []int
	for i, value := range values {
		if i > 0 {
			differences = append(differences, value-values[i-1])
		}
	}
	return differences
}

func (d Day9) isAllZeroes(values []int) bool {
	for _, value := range values {
		if value != 0 {
			return false
		}
	}
	return true
}

func (d Day9) Part2(input string) {
	lines := strings.Split(input, "\r\n")

	var sum int
	for _, line := range lines {
		numberStrings := strings.Split(line, " ")

		var numbers []int
		for _, numberString := range numberStrings {
			number, err := strconv.Atoi(numberString)
			if err != nil {
				panic(err)
			}
			numbers = append(numbers, number)
		}

		reservedNumbers := make([]int, len(numbers))
		for i, number := range numbers {
			reservedNumbers[len(reservedNumbers)-1-i] = number
		}
		numbers = reservedNumbers

		var sequences [][]int
		sequences = append(sequences, numbers)
		for !d.isAllZeroes(sequences[len(sequences)-1]) {
			sequences = append(sequences, d.getDifferences(sequences[len(sequences)-1]))
		}

		sequences[len(sequences)-1] = append(sequences[len(sequences)-1], 0)

		for i := len(sequences) - 2; i > -1; i-- {
			currentSequence := sequences[i]
			previousSequence := sequences[i+1]
			newValue := currentSequence[len(previousSequence)-1] + previousSequence[len(previousSequence)-1]
			sequences[i] = append(sequences[i], newValue)
		}

		sum += sequences[0][len(sequences[0])-1]
	}

	fmt.Println(sum)
}
