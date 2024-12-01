package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Day1Part1() {
	f, err := os.Open("./input/01.txt")
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(f)
	leftList := make([]int, 0)
	rightList := make([]int, 0)
	for s.Scan() {
		line := s.Text()
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

	fmt.Println(diffTotal)
}
