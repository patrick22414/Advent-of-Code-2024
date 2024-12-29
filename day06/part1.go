package day06

import (
	"slices"

	"github.com/patrick22414/Advent-of-Code-2024/readinput"
)

const (
	Empty byte = '.'
	Path  byte = 'X'
	Wall  byte = '#'
	Up    byte = '^'
	Right byte = '>'
	Down  byte = 'v'
	Left  byte = '<'
)

func Part1() int {
	area := make([][]byte, 0)
	x, y := 0, -1 // guard position
	for lineString := range readinput.ReadInput("./input.txt") {
		line := []byte(lineString)
		area = append(area, line)
		if y < 0 {
			y = slices.Index(line, Up) // current position of the guard
			if y < 0 {
				x++
			}
		}
	}

	// starting direction is up
	m, n := len(area), len(area[0])
	dx, dy := -1, 0
	for {
		area[x][y] = Path

		x += dx
		y += dy
		if x < 0 || x >= m || y < 0 || y >= n {
			break
		}

		if area[x][y] == Wall {
			x -= dx
			y -= dy

			// turn right 90 degrees
			d := complex(float64(dx), float64(dy))
			d *= complex(0, -1)
			dx = int(real(d))
			dy = int(imag(d))
		}
	}

	total := 0
	for _, row := range area {
		for _, cell := range row {
			if cell == Path {
				total++
			}
		}
	}
	return total
}
