package day03

import (
	"regexp"
	"strconv"

	"github.com/patrick22414/Advent-of-Code-2024/readinput"
)

var RE_MUL = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

func Part1() int {
	total := 0
	for line := range readinput.ReadInput("./input.txt") {
		exps := RE_MUL.FindAllStringSubmatch(line, -1)
		for _, exp := range exps {
			a, err := strconv.Atoi(exp[1])
			if err != nil {
				panic(err)
			}
			b, err := strconv.Atoi(exp[2])
			if err != nil {
				panic(err)
			}
			total += a * b
		}
	}

	return total
}
