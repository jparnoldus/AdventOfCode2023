package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Day7 struct {
}

type Day7Hand struct {
	Card   string
	Bid    string
	Type   string
	Values []int
}

func (d Day7) translate(card string) []int {
	var answer []int
	for _, label := range card {
		temp := string(label)
		if temp == "T" {
			temp = "10"
		}
		if temp == "J" {
			temp = "11"
		}
		if temp == "Q" {
			temp = "12"
		}
		if temp == "K" {
			temp = "13"
		}
		if temp == "A" {
			temp = "14"
		}

		number, _ := strconv.Atoi(temp)
		answer = append(answer, number)
	}
	return answer
}

func (d Day7) translateWithJokers(card string) []int {
	var answer []int
	for _, label := range card {
		temp := string(label)
		if temp == "J" {
			temp = "1"
		}
		if temp == "T" {
			temp = "10"
		}
		if temp == "Q" {
			temp = "11"
		}
		if temp == "K" {
			temp = "12"
		}
		if temp == "A" {
			temp = "13"
		}

		number, _ := strconv.Atoi(temp)
		answer = append(answer, number)
	}
	return answer
}

func (d Day7) getRepetitions(card []int) map[int]int {
	repetitions := make(map[int]int)
	for _, number := range card {
		repetitions[number]++
	}

	return repetitions
}

func (d Day7) getSorting(repetitions map[int]int) []int {
	keys := make([]int, 0, len(repetitions))
	for key := range repetitions {
		if key != 1 {
			keys = append(keys, key)
		}
	}
	sort.Slice(keys, func(i, j int) bool { return repetitions[keys[i]] > repetitions[keys[j]] })

	return keys
}

func (d Day7) compareCards(card1 []int, card2 []int) bool {
	if card1[0] > card2[0] {
		return true
	} else if card1[0] < card2[0] {
		return false
	}
	if card1[1] > card2[1] {
		return true
	} else if card1[1] < card2[1] {
		return false
	}
	if card1[2] > card2[2] {
		return true
	} else if card1[2] < card2[2] {
		return false
	}
	if card1[3] > card2[3] {
		return true
	} else if card1[3] < card2[3] {
		return false
	}
	if card1[4] > card2[4] {
		return true
	} else if card1[4] < card2[4] {
		return false
	}
	panic("same cards")
}

func (d Day7) Part1(input string) {
	lines := strings.Split(input, "\r\n")

	sort.Slice(lines, func(i, j int) bool {
		line1 := lines[i]
		line2 := lines[j]

		parts := strings.Split(line1, " ")
		card1 := d.translate(parts[0])
		repetitions1 := d.getRepetitions(card1)
		keys1 := d.getSorting(repetitions1)

		parts = strings.Split(line2, " ")
		card2 := d.translate(parts[0])
		repetitions2 := d.getRepetitions(card2)
		keys2 := d.getSorting(repetitions2)

		// No case of having more of a kind and still losing
		if repetitions1[keys1[0]] > repetitions2[keys2[0]] {
			return true
		} else if repetitions1[keys1[0]] < repetitions2[keys2[0]] {
			return false
		}

		// 4 or 5 of a kind
		if repetitions1[keys1[0]] > 3 && repetitions2[keys2[0]] > 3 {
			return d.compareCards(card1, card2)
		}

		// Full house
		if repetitions1[keys1[0]] == 3 && repetitions1[keys1[1]] == 2 && repetitions2[keys2[0]] == 3 && repetitions2[keys2[1]] == 2 {
			return d.compareCards(card1, card2)
		}

		// Three of a kind
		if repetitions1[keys1[0]] == 3 && repetitions2[keys2[0]] == 3 {
			// Sneaky full house
			if repetitions1[keys1[1]] == 2 {
				return true
			} else if repetitions2[keys2[1]] == 2 {
				return false
			}
			return d.compareCards(card1, card2)
		}

		// Two pair
		if repetitions1[keys1[0]] == 2 && repetitions1[keys1[1]] == 2 && repetitions2[keys2[0]] == 2 && repetitions2[keys2[1]] == 2 {
			return d.compareCards(card1, card2)
		}

		// One pair
		if repetitions1[keys1[0]] == 2 && repetitions2[keys2[0]] == 2 {
			if repetitions1[keys1[1]] == 2 {
				return true
			} else if repetitions2[keys2[1]] == 2 {
				return false
			}
			return d.compareCards(card1, card2)
		}

		// High card
		if repetitions1[keys1[0]] == 1 && repetitions2[keys2[0]] == 1 {
			return d.compareCards(card1, card2)
		}

		return false
	})

	var sum int
	for i, line := range lines {
		parts := strings.Split(line, " ")
		bid, _ := strconv.Atoi(parts[1])
		sum += bid * (len(lines) - i)
	}
	fmt.Println(sum)
}

func (d Day7) Part2(input string) {
	lines := strings.Split(input, "\r\n")

	sort.Slice(lines, func(i, j int) bool {
		line1 := lines[i]
		line2 := lines[j]

		parts := strings.Split(line1, " ")
		card1 := d.translateWithJokers(parts[0])
		repetitions1 := d.getRepetitions(card1)
		keys1 := d.getSorting(repetitions1)

		parts = strings.Split(line2, " ")
		card2 := d.translateWithJokers(parts[0])
		repetitions2 := d.getRepetitions(card2)
		keys2 := d.getSorting(repetitions2)

		if len(keys1) == 0 && len(keys2) != 0 {
			if repetitions2[keys2[0]]+repetitions2[1] == 5 {
				return false
			}
			return true
		} else if len(keys1) != 0 && len(keys2) == 0 {
			if repetitions1[keys1[0]]+repetitions1[1] == 5 {
				return true
			}
			return false
		}

		// No case of having more of a kind and still losing
		if repetitions1[keys1[0]]+repetitions1[1] > repetitions2[keys2[0]]+repetitions2[1] {
			return true
		} else if repetitions1[keys1[0]]+repetitions1[1] < repetitions2[keys2[0]]+repetitions2[1] {
			return false
		}

		// 5 of a kind
		if len(keys1) == 1 && len(keys2) == 1 {
			return d.compareCards(card1, card2)
		}

		// 4 of a kind
		if len(keys1) == 2 && repetitions1[keys1[1]] == 1 && len(keys2) == 2 && repetitions2[keys2[1]] == 1 {
			return d.compareCards(card1, card2)
		}

		// Full house
		if len(keys1) == 2 && repetitions1[keys1[1]] >= 2 && len(keys2) == 2 && repetitions2[keys2[1]] >= 2 {
			return d.compareCards(card1, card2)
		}

		if len(keys1) == 2 && repetitions1[keys1[1]] >= 2 {
			return true
		} else if len(keys2) == 2 && repetitions2[keys2[1]] >= 2 {
			return false
		}

		// Three of a kind
		if len(keys1) == 3 && repetitions1[keys1[2]] == 1 && len(keys2) == 3 && repetitions2[keys2[2]] == 1 {
			return d.compareCards(card1, card2)
		}

		// Two pair
		if len(keys1) == 3 && repetitions1[keys1[1]] == 2 && len(keys2) == 3 && repetitions2[keys2[1]] == 2 {
			return d.compareCards(card1, card2)
		}
		if len(keys1) == 3 && repetitions1[keys1[1]] == 2 {
			return true
		} else if len(keys2) == 3 && repetitions2[keys2[1]] == 2 {
			return false
		}

		// One pair
		if (repetitions1[keys1[0]] == 2 || repetitions1[1] == 1) && (repetitions2[keys2[0]] == 2 || repetitions2[1] == 1) {
			return d.compareCards(card1, card2)
		}
		if repetitions1[keys1[0]] == 2 || repetitions1[1] == 1 {
			return true
		} else if repetitions2[keys2[0]] == 2 || repetitions2[1] == 1 {
			return false
		}

		// High card
		if repetitions1[keys1[0]] == 1 && repetitions2[keys2[0]] == 1 {
			return d.compareCards(card1, card2)
		}

		return false
	})

	var sum int
	for i, line := range lines {
		parts := strings.Split(line, " ")
		bid, _ := strconv.Atoi(parts[1])
		sum += bid * (len(lines) - i)
	}
	fmt.Println(sum)
}
