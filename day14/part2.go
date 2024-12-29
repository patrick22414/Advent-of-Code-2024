package day14

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"regexp"

	"github.com/patrick22414/Advent-of-Code-2024/readinput"
)

const S = 7

func RobotsToImage(robots []*Robot) *image.Gray {
	im := image.NewGray(image.Rect(0, 0, W*S, H*S))
	for _, r := range robots {
		for i := range S {
			for j := range S {
				im.Set(r.px*S+i, r.py*S+j, color.White)
			}
		}
	}
	return im
}

var RE_POSSIBLE = regexp.MustCompile(`#{10}`)

func Possible(robots []*Robot) bool {
	area := RobotsToAscii(robots)
	for i := 0; i < W; i++ {
		line := (area[i])[:]
		if RE_POSSIBLE.Match(line) {
			return true
		}
	}
	return false
}

func Part2() int {
	robots := make([]*Robot, 0)
	for line := range readinput.ReadInput("./input.txt") {
		robots = append(robots, NewRobotFromString(line))
	}

	for i := range 10000 {
		if Possible(robots) {
			f, err := os.Create(fmt.Sprintf("%v.png", i))
			if err != nil {
				panic(err)
			}

			png.Encode(f, RobotsToImage(robots))
			return i
		}
		for _, r := range robots {
			r.Move(1)
		}
	}

	return -1 // dummy output
}
