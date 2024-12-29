package day07

import (
	"math"
	"strconv"
	"strings"

	"github.com/patrick22414/Advent-of-Code-2024/readinput"
)

func isEquationPossible3(n int, ns []int, target int) bool {
	if len(ns) == 0 {
		return n == target
	}
	if len(ns) == 1 {
		return n+ns[0] == target ||
			n*ns[0] == target ||
			concat(n, ns[0]) == target
	}
	// then len(ns) > 1

	if n >= target {
		return false
	}

	return isEquationPossible3(n+ns[0], ns[1:], target) ||
		isEquationPossible3(n*ns[0], ns[1:], target) ||
		isEquationPossible3(concat(n, ns[0]), ns[1:], target)
}

func concat(x, y int) int {
	return x*int(math.Pow10(int(math.Log10(float64(y)))+1)) + y
}

func Part2() int {
	total := 0
	for line := range readinput.ReadInput("./input.txt") {
		res, operands, ok := strings.Cut(line, ": ")
		if !ok {
			panic("")
		}
		target, err := strconv.Atoi(res)
		if err != nil {
			panic(err)
		}

		ns := make([]int, 0)
		for _, t := range strings.Split(operands, " ") {
			n, err := strconv.Atoi(t)
			if err != nil {
				panic(err)
			}
			ns = append(ns, n)
		}

		if isEquationPossible3(ns[0], ns[1:], target) {
			// fmt.Println(line, true)
			total += target
		} else {
			// fmt.Println(line, false)
		}
	}

	return total
}
