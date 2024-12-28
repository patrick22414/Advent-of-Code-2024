package day04

import (
	"github.com/patrick22414/Advent-of-Code-2024/readinput"
)

const XMAS = "XMAS"

func Part1() int {
	var m []string
	for line := range readinput.ReadInput("./input.txt") {
		m = append(m, line)
	}

	count := 0
	for i := range m {
		for j := range m[0] {
			if m[i][j] != XMAS[0] {
				continue
			}

			for dx := -1; dx <= 1; dx++ {
				for dy := -1; dy <= 1; dy++ {
					if dx == 0 && dy == 0 {
						continue
					}

					// for each of the 8 directions
					if i+dx*3 < 0 || i+dx*3 >= len(m) {
						continue
					}
					if j+dy*3 < 0 || j+dy*3 >= len(m[0]) {
						continue
					}

					if m[i+dx][j+dy] == XMAS[1] &&
						m[i+dx*2][j+dy*2] == XMAS[2] &&
						m[i+dx*3][j+dy*3] == XMAS[3] {
						count++
					}
				}
			}
		}
	}

	return count
}
