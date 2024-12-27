package day12

import (
	"github.com/patrick22414/Advent-of-Code-2024/grid"
)

type Side struct {
	S grid.Pos // starting corner
	E grid.Pos // ending corner
	D int      // direction
}

const (
	Up = iota
	Right
	Down
	Left
)

type SetOfSides struct {
	sides []Side
}

func NewSetOfSides() *SetOfSides {
	sides := make([]Side, 0)
	return &SetOfSides{sides: sides}
}

func (sos *SetOfSides) Add(s Side) {
	for i, ex := range sos.sides { // for all existing sides
		if ex.D != s.D {
			continue
		}
		if ex.E == s.S {
			sos.sides[i].E = s.E
			return
		}
		if ex.S == s.E {
			sos.sides[i].S = s.S
			return
		}
	}
	sos.sides = append(sos.sides, s)
}

func Part2() int {
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

	areaCounter := make(map[int]int)
	for _, row := range regions {
		for _, r := range row {
			areaCounter[r]++
		}
	}

	sos := make(map[int]*SetOfSides)
	for i := range regionCount {
		sos[i+1] = NewSetOfSides()
	}
	for i := range farm.SizeX() {
		for j := range farm.SizeY() {
			r := regions[i][j]

			topLeft := grid.Pos{X: i, Y: j}
			topRight := grid.Pos{X: i, Y: j + 1}
			bottomRight := grid.Pos{X: i + 1, Y: j + 1}
			bottomLeft := grid.Pos{X: i + 1, Y: j}

			if x := i - 1; x < 0 || regions[x][j] != r {
				sos[r].Add(Side{
					S: topLeft,
					E: topRight,
					D: Right,
				})
			}
			if y := j + 1; y >= farm.SizeY() || regions[i][y] != r {
				sos[r].Add(Side{
					S: topRight,
					E: bottomRight,
					D: Down,
				})
			}
			if x := i + 1; x >= farm.SizeX() || regions[x][j] != r {
				sos[r].Add(Side{
					S: bottomRight,
					E: bottomLeft,
					D: Left,
				})
			}
			if y := j - 1; y < 0 || regions[i][y] != r {
				sos[r].Add(Side{
					S: bottomLeft,
					E: topLeft,
					D: Up,
				})
			}
		}
	}

	// fmt.Println(farm, farm.SizeX())
	// for r, s := range sos {
	// 	fmt.Println(r, len(s.sides))
	// 	for _, side := range s.sides {
	// 		fmt.Println(side)
	// 	}
	// }

	total := 0
	for r, area := range areaCounter {
		total += area * len(sos[r].sides)
	}
	return total
}
