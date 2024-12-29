package day05

import (
	"slices"
	"strconv"
	"strings"

	"github.com/patrick22414/Advent-of-Code-2024/readinput"
)

func Part2() int {
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

	return total
}
