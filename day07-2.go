package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day7Part2(output bool) {
	f, err := os.Open("input/07.txt")
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(f)
	total := 0
	for s.Scan() {
		line := s.Text()
		res, operands, ok := strings.Cut(line, ": ")
		if !ok {
			panic("")
		}
		target, err := strconv.Atoi(res)
		if err != nil {
			panic(err)
		}

		ns := make([]int, 0)
		ts := strings.Split(operands, " ")
		for _, t := range ts {
			n, err := strconv.Atoi(t)
			if err != nil {
				panic(err)
			}
			ns = append(ns, n)
		}

		if isEquationPossible3(ns[0], ns[1:], ts[1:], target) {
			// fmt.Println(line, true)
			total += target
		} else {
			// fmt.Println(line, false)
		}
	}

	if output {
		fmt.Println(total)
	}
}

func isEquationPossible3(n int, ns []int, ts []string, target int) bool {
	if len(ns) == 0 {
		return n == target
	}
	if len(ns) == 1 {
		return n+ns[0] == target ||
			n*ns[0] == target ||
			concat2(strconv.Itoa(n), ts[0]) == target
	}

	// len(ns) > 1
	if n >= target {
		return false
	}

	return isEquationPossible3(n+ns[0], ns[1:], ts[1:], target) ||
		isEquationPossible3(n*ns[0], ns[1:], ts[1:], target) ||
		isEquationPossible3(concat2(strconv.Itoa(n), ts[0]), ns[1:], ts[1:], target)
}

func concat2(x string, y string) int {
	n, err := strconv.Atoi(x + y)
	if err != nil {
		panic("")
	}
	return n
}

func concat(x, y int) int {
	n, err := strconv.Atoi(strconv.Itoa(x) + strconv.Itoa(y))
	if err != nil {
		panic("")
	}
	return n
}
