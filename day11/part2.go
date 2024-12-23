package day11

import (
	"math"
	"strconv"
	"strings"

	"github.com/patrick22414/Advent-of-Code-2024/readinput"
)

func blink(n int) (int, int) {
	if n < 0 {
		panic("n < 0")
	}

	// If the stone is engraved with the number 0, it is replaced by a stone
	// engraved with the number 1.
	if n == 0 {
		return 1, -1
	}

	// If the stone is engraved with a number that has an even number of digits,
	// it is replaced by two stones. The left half of the digits are engraved on
	// the new left stone, and the right half of the digits are engraved on the
	// new right stone. (The new numbers don't keep extra leading zeroes: 1000
	// would become stones 10 and 0.)
	if nDigits := int(math.Log10((float64(n)))) + 1; nDigits%2 == 0 {
		pow10 := int(math.Pow10(nDigits / 2))
		leftValue := n / pow10
		rightValue := n - leftValue*pow10
		return leftValue, rightValue
	}

	// If none of the other rules apply, the stone is replaced by a new stone;
	// the old stone's number multiplied by 2024 is engraved on the new stone.
	return n * 2024, -1
}

func Part2() int {
	stoneCount := make(map[int]int) // stone number and count
	for line := range readinput.ReadInput("./input.txt") {
		for _, s := range strings.Fields(line) {
			v, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			stoneCount[v] += 1
		}
	}

	for range 75 {
		newStoneCount := make(map[int]int, len(stoneCount))
		for n, count := range stoneCount {
			n1, n2 := blink(n)
			newStoneCount[n1] += count
			if n2 >= 0 {
				newStoneCount[n2] += count
			}
		}
		stoneCount = newStoneCount
	}

	total := 0
	for _, count := range stoneCount {
		total += count
	}

	return total
}
