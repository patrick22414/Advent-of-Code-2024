package day08

import "github.com/patrick22414/Advent-of-Code-2024/readinput"

func Part2() int {
	// map of frequency to locations
	antennas := make(map[byte][]Vec)
	i, n := 0, -1
	for line := range readinput.ReadInput("./input.txt") {
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
	isWithinBounds := func(v Vec) bool {
		return v.x >= 0 && v.x < m &&
			v.y >= 0 && v.y < n
	}

	// total := 0
	uniques := make(map[Vec]struct{})
	for _, v := range antennas {
		for i, a1 := range v {
			for _, a2 := range v[i+1:] {
				// fmt.Println(a1, a2)
				d := a2.sub(a1)
				for a := a1; isWithinBounds(a); a = a.sub(d) {
					uniques[a] = struct{}{}
				}
				for a := a2; isWithinBounds(a); a = a.add(d) {
					uniques[a] = struct{}{}
				}
			}
		}
	}

	// fmt.Println(antennas)
	// fmt.Println(uniques)

	return len(uniques)
}
