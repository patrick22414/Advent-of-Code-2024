package day12

import "github.com/patrick22414/Advent-of-Code-2024/map2d"

func Part2() int {
	farm := map2d.FromInput("./input.txt")

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
			pos := map2d.Pos{X: i, Y: j}
			plant := farm.At(pos)

			// depth-first traversal
			stack := []map2d.Pos{pos}
			for len(stack) > 0 {
				// add neighboring, same-plant, not-yet-colored plots
				p := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				regions[p.X][p.Y] = regionCount

				ns := farm.Neighbors4(p)
				for _, n := range ns {
					if farm.At(n) == plant && regions[n.X][n.Y] == 0 {
						stack = append(stack, n)
					}
				}
			}
		}
	}

	// the number of sides in a region is the same as the number of 90-degree
	// angles along its perimeter
	angles := make([][]int, farm.MaxX()) // perimeter of each plot
	for i := range farm.MaxX() {
		angles[i] = make([]int, farm.MaxY())
	}
	// TODO

	areaCounter := make(map[int]int)
	sideCounter := make(map[int]int)
	for i, row := range regions {
		for j, r := range row {
			areaCounter[r]++
			sideCounter[r] += angles[i][j]
		}
	}

	total := 0
	for r, area := range areaCounter {
		total += area * sideCounter[r]
	}
	return total

}
