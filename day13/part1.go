package day13

import (
	"regexp"
	"strconv"

	"github.com/patrick22414/Advent-of-Code-2024/readinput"
)

var RE_BUTTON = regexp.MustCompile(`^Button [AB]: X\+(\d+), Y\+(\d+)$`)
var RE_PRIZE = regexp.MustCompile(`^Prize: X=(\d+), Y=(\d+)$`)

type ClawMachine struct {
	AX, AY, BX, BY int
	PrizeX, PrizeY int
}

func NewClawMachine(ax, ay, bx, by, prizex, prizey string) (*ClawMachine, error) {
	nax, err := strconv.Atoi(ax)
	if err != nil {
		return nil, err
	}
	nay, err := strconv.Atoi(ay)
	if err != nil {
		return nil, err
	}
	nbx, err := strconv.Atoi(bx)
	if err != nil {
		return nil, err
	}
	nby, err := strconv.Atoi(by)
	if err != nil {
		return nil, err
	}
	npx, err := strconv.Atoi(prizex)
	if err != nil {
		return nil, err
	}
	npy, err := strconv.Atoi(prizey)
	if err != nil {
		return nil, err
	}

	return &ClawMachine{
		AX: nax, AY: nay,
		BX: nbx, BY: nby,
		PrizeX: npx, PrizeY: npy,
	}, nil
}

func (cm *ClawMachine) Solve() int {
	minTokens := -1

	for a := 0; ; a++ {
		b := (cm.PrizeX - a*cm.AX) / cm.BX
		if b < 0 {
			break
		}
		if a*cm.AY+b*cm.BY == cm.PrizeY {
			tokens := a*3 + b
			if minTokens < 0 || tokens < minTokens {
				minTokens = tokens
			}
		}
	}

	return minTokens
}

func Part1() int {
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
		cms = append(cms, cm)

		if !ok {
			break
		}
	}

	total := 0
	for _, cm := range cms {
		if tokens := cm.Solve(); tokens > 0 {
			total += tokens
		}
	}

	return total
}
