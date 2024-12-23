package day02

import (
	"strconv"
	"strings"

	"github.com/patrick22414/Advent-of-Code-2024/readinput"
)

func IsSafe(s []int) bool {
	if len(s) == 0 {
		panic("empty slice")
	}

	direction := 0
	for i := 0; i < len(s)-1; i++ {
		a, b := s[i], s[i+1]
		if direction == 0 {
			if a > b {
				// direction is decreasing
				direction = -1
			} else if a < b {
				// direction is increasing
				direction = 1
			}
			// else return false below
		}

		diff := (b - a) * direction
		if diff <= 0 || diff > 3 {
			return false
		}
	}
	return true
}

func Part1() int {
	safeReports := 0
	for line := range readinput.ReadInput("./input.txt") {
		ts := strings.Fields(line)
		ns := make([]int, len(ts))
		for i, token := range ts {
			n, err := strconv.Atoi(token)
			if err != nil {
				panic(err)
			}
			ns[i] = n
		}

		if IsSafe(ns) {
			safeReports++
		}
	}

	return safeReports
}
