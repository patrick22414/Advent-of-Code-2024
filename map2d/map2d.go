package map2d

import (
	"strings"

	"github.com/patrick22414/Advent-of-Code-2024/readinput"
)

type Pos struct {
	X, Y int
}

type Map2D [][]byte

func New() *Map2D {
	m := make(Map2D, 0)
	return &m
}

func FromInput(filename string) *Map2D {
	m := make(Map2D, 0)
	for line := range readinput.ReadInput(filename) {
		m = append(m, []byte(line))
	}
	return &m
}

func (m *Map2D) At(p Pos) byte {
	return (*m)[p.X][p.Y]
}

func (m *Map2D) MaxX() int {
	return len(*m)
}

func (m *Map2D) MaxY() int {
	return len((*m)[0])
}

func (m *Map2D) String() string {
	s := make([]string, 0, len(*m))
	for _, row := range *m {
		s = append(s, string(row))
	}
	return strings.Join(s, "\n")
}

func (m *Map2D) Neighbors2(p Pos) (ns []Pos) {
	if y := p.Y + 1; y < m.MaxY() {
		ns = append(ns, Pos{p.X, y})
	}
	if x := p.X + 1; x < m.MaxX() {
		ns = append(ns, Pos{x, p.Y})
	}
	return
}

func (m *Map2D) Neighbors4(p Pos) (ns []Pos) {
	if x := p.X - 1; x >= 0 {
		ns = append(ns, Pos{x, p.Y})
	}
	if y := p.Y + 1; y < m.MaxY() {
		ns = append(ns, Pos{p.X, y})
	}
	if x := p.X + 1; x < m.MaxX() {
		ns = append(ns, Pos{x, p.Y})
	}
	if y := p.Y - 1; y >= 0 {
		ns = append(ns, Pos{p.X, y})
	}
	return
}

func (m *Map2D) BetterNeighbors4(p Pos, buf *[4]Pos) (ns []Pos) {
	ns = buf[:0]
	if x := p.X - 1; x >= 0 {
		ns = append(ns, Pos{x, p.Y})
	}
	if y := p.Y + 1; y < m.MaxY() {
		ns = append(ns, Pos{p.X, y})
	}
	if x := p.X + 1; x < m.MaxX() {
		ns = append(ns, Pos{x, p.Y})
	}
	if y := p.Y - 1; y >= 0 {
		ns = append(ns, Pos{p.X, y})
	}
	return
}
