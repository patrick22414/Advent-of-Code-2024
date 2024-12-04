package main

import (
	"bufio"
	"fmt"
	"os"
)

func Day4Part2(output bool) {
	f, err := os.Open("input/04.txt")
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(f)
	var m []string
	for s.Scan() {
		line := s.Text()
		m = append(m, line)
	}

	count := 0
	for i := 1; i < len(m)-1; i++ {
		for j := 1; j < len(m[0])-1; j++ {
			if m[i][j] != 'A' {
				continue
			}
			if (m[i-1][j-1] == 'M' && m[i+1][j+1] == 'S') ||
				(m[i-1][j-1] == 'S' && m[i+1][j+1] == 'M') {
				if (m[i-1][j+1] == 'M' && m[i+1][j-1] == 'S') ||
					(m[i-1][j+1] == 'S' && m[i+1][j-1] == 'M') {
					count++
				}
			}
		}
	}

	if output {
		fmt.Println(count)
	}
}
