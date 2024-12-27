package grid

import (
	"strings"

	"github.com/patrick22414/Advent-of-Code-2024/readinput"
)

type Pos struct {
	X, Y int
}

type Grid [][]byte

func New() *Grid {
	g := make(Grid, 0)
	return &g
}

func FromInput(filename string) *Grid {
	g := make(Grid, 0)
	for line := range readinput.ReadInput(filename) {
		g = append(g, []byte(line))
	}
	return &g
}

func (g *Grid) At(p Pos) byte {
	return (*g)[p.X][p.Y]
}

func (g *Grid) SizeX() int {
	return len(*g)
}

func (g *Grid) SizeY() int {
	return len((*g)[0])
}

func (g *Grid) String() string {
	s := make([]string, 0, len(*g))
	for _, row := range *g {
		s = append(s, string(row))
	}
	return strings.Join(s, "\n")
}

func (g *Grid) Neighbors2(p Pos) (ns []Pos) {
	if y := p.Y + 1; y < g.SizeY() {
		ns = append(ns, Pos{p.X, y})
	}
	if x := p.X + 1; x < g.SizeX() {
		ns = append(ns, Pos{x, p.Y})
	}
	return
}

func (g *Grid) Neighbors4(p Pos) (ns []Pos) {
	if x := p.X - 1; x >= 0 {
		ns = append(ns, Pos{x, p.Y})
	}
	if y := p.Y + 1; y < g.SizeY() {
		ns = append(ns, Pos{p.X, y})
	}
	if x := p.X + 1; x < g.SizeX() {
		ns = append(ns, Pos{x, p.Y})
	}
	if y := p.Y - 1; y >= 0 {
		ns = append(ns, Pos{p.X, y})
	}
	return
}

func (g *Grid) BetterNeighbors4(p Pos, buf *[4]Pos) (ns []Pos) {
	ns = buf[:0]
	if x := p.X - 1; x >= 0 {
		ns = append(ns, Pos{x, p.Y})
	}
	if y := p.Y + 1; y < g.SizeY() {
		ns = append(ns, Pos{p.X, y})
	}
	if x := p.X + 1; x < g.SizeX() {
		ns = append(ns, Pos{x, p.Y})
	}
	if y := p.Y - 1; y >= 0 {
		ns = append(ns, Pos{p.X, y})
	}
	return
}
