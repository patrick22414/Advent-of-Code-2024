package day05

import (
	"slices"
	"strconv"
	"strings"

	"github.com/patrick22414/Advent-of-Code-2024/readinput"
)

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

func Part1() int {
	rules := make(map[int][]int, 0)
	input := readinput.ReadInput("./input.txt")
	for line := range input {
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
	for line := range input {
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

	return total
}
