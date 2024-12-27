package day12

import (
	"github.com/patrick22414/Advent-of-Code-2024/grid"
)

func Part1() int {
	farm := grid.FromInput("./input.txt")

	bufNeighbors := [4]grid.Pos{}
	regions, regionCount := make([][]int, farm.SizeX()), 0
	for i := range farm.SizeX() {
		regions[i] = make([]int, farm.SizeY())
	}
	for i := range farm.SizeX() {
		for j := range farm.SizeY() {
			if regions[i][j] != 0 {
				continue
			}

			// create new region
			regionCount++
			pos := grid.Pos{X: i, Y: j}
			plant := farm.At(pos)

			// depth-first traversal
			stack := []grid.Pos{pos}
			for len(stack) > 0 {
				// add neighboring, same-plant, not-yet-colored plots
				p := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				regions[p.X][p.Y] = regionCount

				ns := farm.BetterNeighbors4(p, &bufNeighbors)
				for _, n := range ns {
					if farm.At(n) == plant && regions[n.X][n.Y] == 0 {
						stack = append(stack, n)
					}
				}
			}
		}
	}

	perimeters := make([][]int, farm.SizeX()) // perimeter of each plot
	for i := range farm.SizeX() {
		perimeters[i] = make([]int, farm.SizeY())
		for j := range farm.SizeY() {
			perimeters[i][j] = 4
		}
	}
	for x, row := range *farm {
		for y, plant := range row {
			pos := grid.Pos{X: x, Y: y}
			for _, n := range farm.BetterNeighbors4(pos, &bufNeighbors) {
				if farm.At(n) == plant {
					perimeters[x][y]--
				}
			}
		}
	}

	areaCounter := make(map[int]int)
	periCounter := make(map[int]int)
	for i, row := range regions {
		for j, r := range row {
			areaCounter[r]++
			periCounter[r] += perimeters[i][j]
		}
	}

	total := 0
	for r, area := range areaCounter {
		total += area * periCounter[r]
	}
	return total
}
