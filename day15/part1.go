package day15

import (
	"fmt"
	"slices"

	"github.com/patrick22414/Advent-of-Code-2024/grid"
	"github.com/patrick22414/Advent-of-Code-2024/readinput"
)

const (
	Wall  byte = '#'
	Robot byte = '@'
	Box   byte = 'O'
	Empty byte = '.'
)

const (
	Up    byte = '^'
	Down  byte = 'v'
	Left  byte = '<'
	Right byte = '>'
)

type Direction struct {
	dx, dy int
}

func ToDirection(d byte) Direction {
	switch d {
	case Up:
		return Direction{-1, 0}
	case Right:
		return Direction{0, 1}
	case Down:
		return Direction{1, 0}
	case Left:
		return Direction{0, -1}
	default:
		panic(fmt.Sprintf("not a direction: %+v", (d)))
	}
}

func Move(g *grid.Grid, p grid.Pos, d Direction) grid.Pos {
	p1 := grid.Pos{X: p.X + d.dx, Y: p.Y + d.dy}

	switch g.At(p1) {
	case Wall:
		return p
	case Box:
		p2 := Move(g, p1, d)
		if p2 == p1 {
			return p
		}
		g.Set(p1, g.At(p))
		g.Set(p, Empty)
		return p1
	case Empty:
		g.Set(p1, g.At(p))
		g.Set(p, Empty)
		return p1
	default:
		panic("what?")
	}
}

func Part1() int {
	i, input := 0, readinput.ReadBytes("./input.txt")
	robot := grid.Pos{}
	g := grid.New()
	for line := range input {
		if len(line) == 0 {
			break
		}

		j := slices.Index(line, Robot)
		if j > 0 {
			robot = grid.Pos{X: i, Y: j}
		}

		i++
		g.AppendLine(slices.Clone(line))
	}
	moves := make([]byte, 0, 20000)
	for line := range input {
		moves = append(moves, line...)
	}

	for _, m := range moves {
		d := ToDirection(m)
		robot = Move(g, robot, d)
	}

	total := 0
	for i := range g.SizeX() {
		for j := range g.SizeY() {
			if g.At(grid.Pos{X: i, Y: j}) == Box {
				total += 100*i + j
			}
		}
	}

	return total
}
