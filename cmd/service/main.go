package main

import (
	"os"
)

type Day interface {
	Part1(input string)
	Part2(input string)
}

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	dat, err := os.ReadFile(pwd + "/input/day12.txt")
	if err != nil {
		panic(err)
	}

	day := Day12{}
	day.Part2(string(dat))
}
