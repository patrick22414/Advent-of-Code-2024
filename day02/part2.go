package day02

import (
	"strconv"
	"strings"

	"github.com/patrick22414/Advent-of-Code-2024/readinput"
)

func IsSafeDampened(s []int) bool {
	if IsSafe(s) {
		return true
	}
	// fmt.Printf("\n%v not safe\n", s)
	for i := range s {
		t := make([]int, i, len(s)-1)
		copy(t, s[:i])
		t = append(t, s[i+1:]...)
		if IsSafe(t) {
			// fmt.Printf("removing %v at %v makes it safe: %v\n", s[i], i, t)
			return true
		}
	}
	// fmt.Println("still not safe")
	return false
}

func Part2() int {
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

		if IsSafeDampened(ns) {
			safeReports++
		}
	}

	return safeReports
}
