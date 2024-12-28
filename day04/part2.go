package day04

import (
	"github.com/patrick22414/Advent-of-Code-2024/readinput"
)

func Part2() int {
	var m []string
	for line := range readinput.ReadInput("./input.txt") {
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

	return count
}
