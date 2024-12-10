package main

import (
	"bufio"
	"fmt"
	"os"
)

type Vec struct{ x, y int }

func (v *Vec) add(other Vec) Vec {
	return Vec{v.x + other.x, v.y + other.y}
}
func (v *Vec) sub(other Vec) Vec {
	return Vec{v.x - other.x, v.y - other.y}
}

func antinodes(a1, a2 Vec) []Vec {
	d := a2.sub(a1)
	return []Vec{
		a2.add(d),
		a1.sub(d),
	}
}

func Day8Part1(output bool) {
	f, err := os.Open("input/08.txt")
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(f)
	// map of frequency to locations
	antennas := make(map[byte][]Vec)
	i, n := 0, -1
	for s.Scan() {
		line := s.Text()
		if n == -1 {
			n = len(line) // max y of input
		}
		for j := 0; j < len(line); j++ {
			f := line[j]
			if f != '.' {
				antennas[f] = append(antennas[f], Vec{i, j})
			}
		}
		i++
	}
	m := i // max x of input

	uniques := make(map[Vec]struct{})
	for _, v := range antennas {
		for i, a1 := range v {
			for _, a2 := range v[i+1:] {
				antis := antinodes(a1, a2)
				for _, anti := range antis {
					uniques[anti] = struct{}{}
				}
			}
		}
	}

	total := 0
	for anti := range uniques {
		if anti.x >= 0 && anti.x < m &&
			anti.y >= 0 && anti.y < n {
			total++
		}
	}
	if output {
		fmt.Println(total)
	}
}
