package main

import (
	"bufio"
	"fmt"
	"math"
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

	if output {
		fmt.Println(total)
	}
}

func isEquationPossible3(n int, ns []int, target int) bool {
	if len(ns) == 0 {
		return n == target
	}
	if len(ns) == 1 {
		return n+ns[0] == target ||
			n*ns[0] == target ||
			concat(n, ns[0]) == target
	}

	// len(ns) > 1
	if n >= target {
		return false
	}

	return isEquationPossible3(n+ns[0], ns[1:], target) ||
		isEquationPossible3(n*ns[0], ns[1:], target) ||
		isEquationPossible3(concat(n, ns[0]), ns[1:], target)
}

func concat(x, y int) int {
	return x*int(math.Pow10(int(math.Log10(float64(y))))) + y
}
