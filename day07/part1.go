package day07

import (
	"strconv"
	"strings"

	"github.com/patrick22414/Advent-of-Code-2024/readinput"
)

func isEquationPossible(n int, ns []int, target int) bool {
	if len(ns) == 0 {
		return n == target
	}
	if len(ns) == 1 {
		return n+ns[0] == target || n*ns[0] == target
	}
	// then len(ns) > 1

	if n >= target {
		return false
	}

	return isEquationPossible(n+ns[0], ns[1:], target) || isEquationPossible(n*ns[0], ns[1:], target)
}

func Part1() int {
	total := 0
	for line := range readinput.ReadInput("./input.txt") {
		res, ts, ok := strings.Cut(line, ": ")
		if !ok {
			panic("")
		}
		target, err := strconv.Atoi(res)
		if err != nil {
			panic(err)
		}

		ns := make([]int, 0)
		for _, t := range strings.Split(ts, " ") {
			n, err := strconv.Atoi(t)
			if err != nil {
				panic(err)
			}
			ns = append(ns, n)
		}

		if isEquationPossible(ns[0], ns[1:], target) {
			// fmt.Println(line, true)
			total += target
		} else {
			// fmt.Println(line, false)
		}
	}

	return total
}
