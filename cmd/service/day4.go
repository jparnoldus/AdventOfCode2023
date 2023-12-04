package main

import (
	"fmt"
	"strings"
)

type Day4 struct {
}

func (d *Day4) Part1(input string) {
	lines := strings.Split(input, "\r\n")

	var total int
	for _, line := range lines {
		card := strings.Split(line, ":")
		numbers := strings.Split(card[1], "|")
		winning := strings.Split(strings.TrimSpace(numbers[0]), " ")
		check := strings.Split(strings.TrimSpace(numbers[1]), " ")

		var points int
		for _, s1 := range check {
			if s1 == "" {
				continue
			}
			for _, s2 := range winning {
				if s2 == "" {
					continue
				}

				if s1 == s2 {
					if points == 0 {
						points = 1
					} else {
						points *= 2
					}
				}
			}
		}
		total += points
	}

	fmt.Println(total)
}

func (d *Day4) Part2(input string) {
	lines := strings.Split(input, "\r\n")

	var cards [][][]string
	for _, line := range lines {
		card := strings.Split(line, ":")
		numbers := strings.Split(card[1], "|")
		winning := strings.Split(strings.TrimSpace(numbers[0]), " ")
		check := strings.Split(strings.TrimSpace(numbers[1]), " ")

		cards = append(cards, [][]string{
			winning,
			check,
		})
	}

	counter := make([]int, len(cards))
	for i, card := range cards {
		counter[i]++

		var amount int
		for _, s1 := range card[0] {
			if s1 == "" {
				continue
			}
			for _, s2 := range card[1] {
				if s2 == "" {
					continue
				}

				if s1 == s2 {
					amount++
				}
			}
		}

		for j := i + 1; j < i+1+amount; j++ {
			if j >= len(cards) {
				continue
			}

			counter[j] += counter[i]
		}
	}

	var sum int
	for _, count := range counter {
		sum += count
	}

	fmt.Println(sum)
}
