package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Day5Part2(output bool) {
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
		book2 := slices.Clone(book)
		slices.SortFunc(book2, func(a, b int) int {
			if slices.Contains(rules[a], b) {
				return -1 // a < b
			}
			return 0
		})

		if !slices.Equal(book, book2) {
			total += book2[len(book2)/2]
		}
	}

	if output {
		fmt.Println(total)
	}
}
