package day15

import (
	"github.com/patrick22414/Advent-of-Code-2024/grid"
	"github.com/patrick22414/Advent-of-Code-2024/readinput"
)

const (
	BoxL byte = '['
	BoxR byte = ']'
)

func CanMoveWide(g *grid.Grid, p grid.Pos, d Direction) bool {
	p1 := grid.Pos{X: p.X + d.dx, Y: p.Y + d.dy}

	switch g.At(p1) {
	case Wall:
		return false
	case BoxL:
		other := grid.Pos{X: p1.X, Y: p1.Y + 1}
		if d == ToDirection(Right) {
			return CanMoveWide(g, other, d)
		}
		return CanMoveWide(g, p1, d) && CanMoveWide(g, other, d)
	case BoxR:
		other := grid.Pos{X: p1.X, Y: p1.Y - 1}
		if d == ToDirection(Left) {
			return CanMoveWide(g, other, d)
		}
		return CanMoveWide(g, p1, d) && CanMoveWide(g, other, d)
	case Empty:
		return true
	default:
		panic("what?")
	}
}

func MoveWide(g *grid.Grid, p grid.Pos, d Direction) grid.Pos {
	p1 := grid.Pos{X: p.X + d.dx, Y: p.Y + d.dy}

	switch g.At(p1) {
	case BoxL:
		other := grid.Pos{X: p1.X, Y: p1.Y + 1}
		if d == ToDirection(Right) {
			MoveWide(g, other, d)
			g.Set(other, BoxL)
		} else {
			MoveWide(g, p1, d)
			MoveWide(g, other, d)
		}
	case BoxR:
		other := grid.Pos{X: p1.X, Y: p1.Y - 1}
		if d == ToDirection(Left) {
			MoveWide(g, other, d)
			g.Set(other, BoxR)
		} else {
			MoveWide(g, p1, d)
			MoveWide(g, other, d)
		}
	case Wall:
		panic("hit a wall!")
	}

	g.Set(p1, g.At(p))
	g.Set(p, Empty)

	return p1
}

func Part2() int {
	i, input := 0, readinput.ReadBytes("./input.txt")
	robot := grid.Pos{}
	g := grid.New()
	for line := range input {
		if len(line) == 0 {
			break
		}

		lineWide := make([]byte, 0, len(line)*2)
		for j, b := range line {
			switch b {
			case Wall:
				lineWide = append(lineWide, Wall, Wall)
			case Box:
				lineWide = append(lineWide, BoxL, BoxR)
			case Empty:
				lineWide = append(lineWide, Empty, Empty)
			case Robot:
				robot = grid.Pos{X: i, Y: j * 2}
				lineWide = append(lineWide, Robot, Empty)
			}
		}

		i++
		g.AppendLine(lineWide)
	}
	moves := make([]byte, 0, 20000)
	for line := range input {
		moves = append(moves, line...)
	}

	for _, m := range moves {
		d := ToDirection(m)
		if CanMoveWide(g, robot, d) {
			robot = MoveWide(g, robot, d)
		}
	}

	total := 0
	for i := range g.SizeX() {
		for j := range g.SizeY() {
			if g.At(grid.Pos{X: i, Y: j}) == BoxL {
				total += 100*i + j
			}
		}
	}

	return total
}
