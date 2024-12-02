package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day2Part2() {
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

		if IsSafeDampened(ns) {
			safeReports++
		}
	}

	fmt.Println(safeReports)
}

func IsSafeDampened(s []int) bool {
	if IsSafe(s) {
		return true
	}
	fmt.Printf("\n%v not safe\n", s)
	for i := range s {
		t := make([]int, i, len(s)-1)
		copy(t, s[:i])
		t = append(t, s[i+1:]...)
		if IsSafe(t) {
			fmt.Printf("removing %v at %v makes it safe: %v\n", s[i], i, t)
			return true
		}
	}
	fmt.Println("still not safe")
	return false
}
