package day14

import (
	"fmt"
	"image/color"
	"strconv"
	"strings"

	"github.com/patrick22414/Advent-of-Code-2024/readinput"
)

const (
	// W = 11
	// H = 7

	W = 101
	H = 103
)

func MustAtoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

type Robot struct {
	px, py int
	vx, vy int
	c      color.Color
}

func NewRobotFromString(line string) *Robot {
	fields := strings.Fields(line)
	px, py, _ := strings.Cut(fields[0][len("p="):], ",")
	vx, vy, _ := strings.Cut(fields[1][len("v="):], ",")

	return &Robot{
		px: MustAtoi(px),
		py: MustAtoi(py),
		vx: MustAtoi(vx),
		vy: MustAtoi(vy),
	}
}

func (r *Robot) String() string {
	return fmt.Sprintf("p=%v,%v v=%v,%v", r.px, r.py, r.vx, r.vy)
}

func (r *Robot) Move(seconds int) {
	r.px += r.vx * seconds
	r.py += r.vy * seconds
	r.px %= W
	r.py %= H
	if r.px < 0 {
		r.px += W
	}
	if r.py < 0 {
		r.py += H
	}
}

func (r *Robot) Quadrant() int {
	if r.px > W/2 {
		if r.py > H/2 {
			return 4
		} else if r.py < H/2 {
			return 1
		} else {
			return 0
		}
	} else if r.px < W/2 {
		if r.py > H/2 {
			return 3
		} else if r.py < H/2 {
			return 2
		} else {
			return 0
		}
	}
	return 0
}

var AsciiArt [H][W]byte

func RobotsToAscii(robots []*Robot) *[H][W]byte {
	for i := range H {
		for j := range W {
			AsciiArt[i][j] = '.'
		}
	}
	for _, r := range robots {
		AsciiArt[r.py][r.px] = '#'
	}
	return &AsciiArt
}

func PrintArea(robots []*Robot) string {
	var area [H][W]byte
	for i := range H {
		for j := range W {
			area[i][j] = '.'
		}
	}
	for _, r := range robots {
		if area[r.py][r.px] == '.' {
			area[r.py][r.px] = '1'
		} else {
			area[r.py][r.px]++
		}
	}

	s := ""
	for _, line := range area {
		s += string(line[:]) + "\n"
	}
	return s
}

func Part1() int {
	robots := make([]*Robot, 0)
	for line := range readinput.ReadInput("./input.txt") {
		robots = append(robots, NewRobotFromString(line))
	}

	// fmt.Println(PrintArea(robots))

	quands := make(map[int]int, 4)
	for _, r := range robots {
		r.Move(100)
		if q := r.Quadrant(); q != 0 {
			quands[q]++
		}
	}

	// fmt.Println(PrintArea(robots))

	answer := 1
	for _, v := range quands {
		answer *= v
	}

	return answer
}
