package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Day5Part1(output bool) {
	f, err := os.Open("input/05.txt")
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(f)

	// scan for rules
	rules := make(map[int][]int, 0)
	for s.Scan() {
		line := s.Text()
		a, b, ok := strings.Cut(line, "|")
		if !ok {
			break
		}
		x, err := strconv.Atoi(a)
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(b)
		if err != nil {
			panic(err)
		}

		// rule, ok := rules[x]
		rules[x] = append(rules[x], y)
	}

	// scan for books (sequences of pages)
	books := make([][]int, 0)
	for s.Scan() {
		line := s.Text()
		pages := strings.Split(line, ",")
		book := make([]int, 0)
		for _, p := range pages {
			n, err := strconv.Atoi(p)
			if err != nil {
				panic(err)
			}
			book = append(book, n)
		}
		books = append(books, book)
	}

	total := 0
	for _, book := range books {
		if IsOrderCorrect(book, rules) {
			total += book[len(book)/2]
		}
	}

	if output {
		fmt.Println(total)
	}
}

func IsOrderCorrect(book []int, rules map[int][]int) bool {
	for j, b := range book {
		rule, ok := rules[b]
		if !ok {
			continue
		}

		for _, a := range book[:j] {
			if slices.Contains(rule, a) {
				return false
			}
		}
	}
	return true
}
