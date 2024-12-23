package day01

import (
	"strconv"
	"strings"

	"github.com/patrick22414/Advent-of-Code-2024/readinput"
)

func Part2() int {
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

	// frequency of numbers in the right list
	freq := make(map[int]int)
	for i := 0; i < len(rightList); i++ {
		n, ok := freq[rightList[i]]
		if !ok {
			freq[rightList[i]] = 1
		} else {
			freq[rightList[i]] = n + 1
		}
	}

	similarityScore := 0
	for i := 0; i < len(leftList); i++ {
		n, ok := freq[leftList[i]]
		if ok {
			similarityScore += leftList[i] * n
		}
	}

	return similarityScore
}
