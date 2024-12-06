package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func Day6Part2(output bool) {
	f, err := os.Open("input/06.txt")
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(f)

	area := make([][]byte, 0)
	x, y := 0, -1 // guard position
	for s.Scan() {
		line := []byte(s.Text())
		area = append(area, line)
		if y < 0 {
			y = slices.Index(line, byte(Up)) // current position of the guard
			if y < 0 {
				x++
			}
		}
	}

	areaClean := slices.Clone(area)
	areaCopy := slices.Clone(area)
	for i, row := range area {
		areaClean[i] = slices.Clone(row)
		areaCopy[i] = slices.Clone(row)
	}

	// mark the area
	Patrol(area, x, y)
	// total := 0
	// for _, row := range area {
	// 	fmt.Println(string(row))
	// 	for _, cell := range row {
	// 		if cell != Empty && cell != Wall {
	// 			total++
	// 		}
	// 	}
	// }

	// then try putting Wall at each marked position
	count := 1
	countLoops := 0
	for i, row := range area {
		for j, cell := range row {
			if i == x && j == y {
				continue // cannot put Wall at Guard position
			}

			if cell != Empty && cell != Wall {
				// fmt.Println(count, "/", total-1) // progress bar
				count++

				// clear out test area
				for ii, rowClean := range areaClean {
					copy(areaCopy[ii], rowClean)
				}
				areaCopy[i][j] = Wall // add a new Wall
				if Patrol(areaCopy, x, y) {
					countLoops++
				}
			}
		}
	}

	if output {
		fmt.Println(countLoops)
	}
}

func DirectionMarker(dx, dy int) byte {
	if dx == -1 && dy == 0 {
		return Up
	} else if dx == 0 && dy == 1 {
		return Right
	} else if dx == 1 && dy == 0 {
		return Down
	} else if dx == 0 && dy == -1 {
		return Left
	}
	panic("not a direction")
}

func MarkerDirection(marker byte) (dx, dy int) {
	if marker == Up {
		dx, dy = -1, 0
	} else if marker == Right {
		dx, dy = 0, 1
	} else if marker == Down {
		dx, dy = 1, 0
	} else if marker == Left {
		dx, dy = 0, -1
	} else {
		panic("not a starting position")
	}
	return
}

// Given a map of the area and a starting position (x, y), mark the path of the
// patrolling guard. Return if the guard is stuck in a loop (true) or will exit
// the area (false).
func Patrol(area [][]byte, x, y int) (isLoop bool) {
	// the starting direction is already marked
	dx, dy := MarkerDirection(area[x][y])
	area[x][y] = Empty

	m, n := len(area), len(area[0]) // size of the area
	marker := DirectionMarker(dx, dy)
	for {
		x1, y1 := x+dx, y+dy
		if x1 < 0 || x1 >= m || y1 < 0 || y1 >= n {
			// out of bounds
			area[x][y] = marker
			break
		}

		if area[x1][y1] == Wall {
			// turn right 90 degrees
			d := complex(float64(dx), float64(dy))
			d *= complex(0, -1)
			dx = int(real(d))
			dy = int(imag(d))
			marker = DirectionMarker(dx, dy)
		} else {
			if area[x][y] == marker {
				return true
			}
			area[x][y] = marker
			x, y = x1, y1
		}
	}
	return false
}
