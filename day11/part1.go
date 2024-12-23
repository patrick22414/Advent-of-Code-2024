package day11

import (
	"math"
	"strconv"
	"strings"

	"github.com/patrick22414/Advent-of-Code-2024/readinput"
)

type Stone struct {
	Value int // number engraved on the stone; -1 for non-left nodes
	Left  *Stone
	Right *Stone
}

func (s *Stone) Count() int {
	if s.Value >= 0 {
		return 1
	}
	return s.Left.Count() + s.Right.Count()
}

func (s *Stone) Blink() {
	if s.Value < 0 {
		s.Left.Blink()
		s.Right.Blink()
		return
	}

	// If the stone is engraved with the number 0, it is replaced by a stone
	// engraved with the number 1.
	if s.Value == 0 {
		s.Value = 1
		return
	}

	// If the stone is engraved with a number that has an even number of digits,
	// it is replaced by two stones. The left half of the digits are engraved on
	// the new left stone, and the right half of the digits are engraved on the
	// new right stone. (The new numbers don't keep extra leading zeroes: 1000
	// would become stones 10 and 0.)
	if nDigits := int(math.Log10((float64(s.Value)))) + 1; nDigits%2 == 0 {
		pow10 := int(math.Pow10(nDigits / 2))
		leftValue := s.Value / pow10
		rightValue := s.Value - leftValue*pow10
		s.Left = &Stone{leftValue, nil, nil}
		s.Right = &Stone{rightValue, nil, nil}
		s.Value = -1
		return
	}

	// If none of the other rules apply, the stone is replaced by a new stone;
	// the old stone's number multiplied by 2024 is engraved on the new stone.
	s.Value *= 2024
	return
}

func Part1() int {
	stones := make([]*Stone, 0)
	for line := range readinput.ReadInput("./input.txt") {
		for _, s := range strings.Fields(line) {
			v, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			stones = append(stones, &Stone{v, nil, nil})
		}
	}

	for range 25 {
		for _, stone := range stones {
			stone.Blink()
		}
	}

	total := 0
	for _, stone := range stones {
		// fmt.Println(stone)
		total += stone.Count()
	}

	return total
}
