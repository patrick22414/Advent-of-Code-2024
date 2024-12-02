package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day2Part1() {
	f, err := os.Open("./input/02.txt")
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(f)
	safeReports := 0
	for s.Scan() {
		line := s.Text()

		ts := strings.Split(line, " ")
		ns := make([]int, len(ts))
		for i, token := range ts {
			ns[i], err = strconv.Atoi(token)
			if err != nil {
				panic(err)
			}
		}

		if IsSafe(ns) {
			safeReports++
		}
	}

	fmt.Println(safeReports)
}

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
