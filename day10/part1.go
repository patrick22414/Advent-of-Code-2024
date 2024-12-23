package day10

import (
	"github.com/patrick22414/Advent-of-Code-2024/readinput"
)

type Pos struct {
	x, y int
}

type Map2D [][]byte

func (m *Map2D) At(p Pos) byte {
	return (*m)[p.x][p.y]
}

func (m *Map2D) Neighbors4(p Pos) (ns []Pos) {
	if x := p.x - 1; x >= 0 {
		ns = append(ns, Pos{x, p.y})
	}
	if y := p.y + 1; y < len((*m)[0]) {
		ns = append(ns, Pos{p.x, y})
	}
	if x := p.x + 1; x < len(*m) {
		ns = append(ns, Pos{x, p.y})
	}
	if y := p.y - 1; y >= 0 {
		ns = append(ns, Pos{p.x, y})
	}
	return
}

func (m *Map2D) Score(trailhead Pos) (score int) {
	if m.At(trailhead) != '0' {
		panic("not a trailhead")
	}

	ps := []Pos{trailhead}
	summits := make(map[Pos]struct{})
	for len(ps) > 0 {
		p := ps[len(ps)-1]
		ps = ps[:len(ps)-1] // pop

		h := m.At(p)
		if h == '9' {
			summits[p] = struct{}{}
			continue
		}

		ns := m.Neighbors4(p)
		for _, n := range ns {
			if m.At(n) == h+1 {
				ps = append(ps, n) // push
			}
		}
	}
	return len(summits)
}

func Part1() int {
	m, i := Map2D{}, 0
	trailheads := make([]Pos, 0)
	for line := range readinput.ReadInput("./input.txt") {
		m = append(m, []byte(line))
		for j := 0; j < len(line); j++ {
			if line[j] == '0' {
				trailheads = append(trailheads, Pos{i, j})
			}
		}
		i++
	}

	// for _, v := range m {
	// 	fmt.Println(string(v))
	// }

	answer := 0
	for _, h := range trailheads {
		// fmt.Println("trailhead", h, "score:", m.Score(h))
		answer += m.Score(h)
	}

	return answer
}
