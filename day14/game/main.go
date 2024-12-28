package main

import (
	"fmt"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/patrick22414/Advent-of-Code-2024/day14"
	"github.com/patrick22414/Advent-of-Code-2024/readinput"
)

type Game struct {
	robots  []*day14.Robot
	img     *image.Gray
	seconds int
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		for range 1000 {
			for _, r := range g.robots {
				r.Move(1)
			}
			g.seconds++
			if day14.Possible(g.robots) {
				break
			}
		}

		g.img = day14.RobotsToImage(g.robots)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(ebiten.NewImageFromImage(g.img), nil)
	ebitenutil.DebugPrint(
		screen,
		fmt.Sprintf("Press SPACE to search over the next 1000 states until you find the Christmas tree\nCurrent:%v", g.seconds),
	)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

const (
	ScreenWidth  = day14.W * day14.S
	ScreenHeight = day14.H * day14.S
)

func main() {
	robots := make([]*day14.Robot, 0)
	for line := range readinput.ReadInput("./input.txt") {
		robots = append(robots, day14.NewRobotFromString(line))
	}

	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Day 14 Part 2")
	ebiten.SetTPS(60)
	if err := ebiten.RunGame(&Game{
		robots: robots,
		img:    day14.RobotsToImage(robots),
	}); err != nil {
		log.Fatal(err)
	}
}
