package day13

import (
	"github.com/patrick22414/Advent-of-Code-2024/readinput"
)

func (cm *ClawMachine) Solve2() int {
	det := cm.AX*cm.BY - cm.AY*cm.BX
	if det == 0 {
		return -1
	}

	a := (cm.BY*cm.PrizeX - cm.BX*cm.PrizeY) / det
	b := (cm.AX*cm.PrizeY - cm.AY*cm.PrizeX) / det
	if a < 0 || b < 0 {
		return -1
	}
	if cm.AX*a+cm.BX*b != cm.PrizeX || cm.AY*a+cm.BY*b != cm.PrizeY {
		return -1
	}

	minTokens := 3*a + b
	return minTokens
}

func Part2() int {
	input := readinput.ReadInput("./input.txt")
	cms := make([]*ClawMachine, 0)
	for {
		lineA := <-input
		lineB := <-input
		linePrize := <-input
		_, ok := <-input // empty line

		ma := RE_BUTTON.FindStringSubmatch(lineA)
		mb := RE_BUTTON.FindStringSubmatch(lineB)
		mp := RE_PRIZE.FindStringSubmatch(linePrize)

		cm, err := NewClawMachine(
			ma[1], ma[2],
			mb[1], mb[2],
			mp[1], mp[2],
		)
		if err != nil {
			panic(err)
		}
		cm.PrizeX += 10000000000000
		cm.PrizeY += 10000000000000
		cms = append(cms, cm)

		if !ok {
			break
		}
	}

	total := 0
	for _, cm := range cms {
		if tokens := cm.Solve2(); tokens > 0 {
			total += tokens
		}
	}

	return total
}
