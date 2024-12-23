package day01

import (
	"slices"
	"strconv"
	"strings"

	"github.com/patrick22414/Advent-of-Code-2024/readinput"
)

func Part1() int {
	leftList := make([]int, 0)
	rightList := make([]int, 0)
	for line := range readinput.ReadInput("./input.txt") {
		ls, rs, ok := strings.Cut(line, "   ")
		if !ok {
			continue
		}

		ln, err := strconv.Atoi(ls)
		if err != nil {
			panic(err)
		}
		leftList = append(leftList, ln)

		rn, err := strconv.Atoi(rs)
		if err != nil {
			panic(err)
		}
		rightList = append(rightList, rn)
	}

	slices.Sort(leftList)
	slices.Sort(rightList)

	diffTotal := 0
	for i := 0; i < len(leftList); i++ {
		d := leftList[i] - rightList[i]
		if d >= 0 {
			diffTotal += d
		} else {
			diffTotal -= d
		}
	}

	return diffTotal
}
