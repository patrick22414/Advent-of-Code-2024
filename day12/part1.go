package day12

import (
	"github.com/patrick22414/Advent-of-Code-2024/grid"
)

func Part1() int {
	farm := grid.FromInput("./input.txt")

	bufNeighbors := [4]grid.Pos{}
	// fmt.Println(bufNeighbors)
	regions, regionCount := make([][]int, farm.MaxX()), 0
	for i := range farm.MaxX() {
		regions[i] = make([]int, farm.MaxY())
	}
	for i := range farm.MaxX() {
		for j := range farm.MaxY() {
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

	perimeters := make([][]int, farm.MaxX()) // perimeter of each plot
	for i := range farm.MaxX() {
		perimeters[i] = make([]int, farm.MaxY())
		for j := range farm.MaxY() {
			perimeters[i][j] = 4
		}
	}
	for x, row := range *farm {
		for y, plant := range row {
			pos := grid.Pos{X: x, Y: y}
			for _, n := range farm.Neighbors4(pos) {
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
