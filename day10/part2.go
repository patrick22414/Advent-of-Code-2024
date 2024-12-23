package day10

import (
	"github.com/patrick22414/Advent-of-Code-2024/readinput"
)

func (m *Map2D) Rating(trailhead Pos) (rating int) {
	if m.At(trailhead) != '0' {
		panic("not a trailhead")
	}

	ps := []Pos{trailhead}
	for len(ps) > 0 {
		p := ps[len(ps)-1]
		ps = ps[:len(ps)-1] // pop

		h := m.At(p)
		if h == '9' {
			rating++
			continue
		}

		ns := m.Neighbors4(p)
		for _, n := range ns {
			if m.At(n) == h+1 {
				ps = append(ps, n) // push
			}
		}
	}
	return
}

func Part2() int {
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

	answer := 0
	for _, h := range trailheads {
		answer += m.Rating(h)
	}

	return answer
}
