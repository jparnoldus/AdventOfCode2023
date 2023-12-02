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

	dat, err := os.ReadFile(pwd + "/input/day2.txt")
	if err != nil {
		panic(err)
	}

	day := Day2{}
	day.Part2(string(dat))
}
