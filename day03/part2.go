package day03

import (
	"cmp"
	"regexp"
	"slices"
	"strconv"

	"github.com/patrick22414/Advent-of-Code-2024/readinput"
)

var RE_DO = regexp.MustCompile(`do\(\)`)
var RE_DONT = regexp.MustCompile(`don't\(\)`)

type Instruction int

const (
	Do Instruction = iota
	Dont
	Mul
)

type Match struct {
	loc  int
	ins  Instruction
	a, b string
}

func Part2() int {
	enabled := true
	total := 0
	for line := range readinput.ReadInput("input.txt") {
		dos := RE_DO.FindAllStringSubmatchIndex(line, -1)
		donts := RE_DONT.FindAllStringSubmatchIndex(line, -1)
		muls := RE_MUL.FindAllStringSubmatchIndex(line, -1)

		matches := make([]Match, 0, len(dos)+len(donts)+len(muls))
		for _, do := range dos {
			matches = append(matches, Match{
				loc: do[0],
				ins: Do,
			})
		}
		for _, dont := range donts {
			matches = append(matches, Match{
				loc: dont[0],
				ins: Dont,
			})
		}
		for _, mul := range muls {
			matches = append(matches, Match{
				loc: mul[0],
				ins: Mul,
				a:   line[mul[2]:mul[3]],
				b:   line[mul[4]:mul[5]],
			})
		}

		slices.SortFunc(matches, func(x, y Match) int {
			return cmp.Compare(x.loc, y.loc)
		})

		for _, m := range matches {
			if m.ins == Do {
				enabled = true
			} else if m.ins == Dont {
				enabled = false
			} else if m.ins == Mul {
				if !enabled {
					continue
				}
				a, err := strconv.Atoi(m.a)
				if err != nil {
					panic(err)
				}
				b, err := strconv.Atoi(m.b)
				if err != nil {
					panic(err)
				}
				total += a * b
			}
		}
	}
	return total
}
