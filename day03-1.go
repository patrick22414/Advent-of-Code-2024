package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var RE_MUL = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

func Day3Part1() {
	f, err := os.Open("input/03.txt")
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(f)
	total := 0
	for s.Scan() {
		line := s.Text()
		exps := RE_MUL.FindAllStringSubmatch(line, -1)
		for _, exp := range exps {
			a, err := strconv.Atoi(exp[1])
			if err != nil {
				panic(err)
			}
			b, err := strconv.Atoi(exp[2])
			if err != nil {
				panic(err)
			}
			total += a * b
		}
	}

	fmt.Println(total)
}
