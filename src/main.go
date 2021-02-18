package main

import (
	"log"

	"github.com/danieldhoang93/ebiten-shooter/objects"
	"github.com/hajimehoshi/ebiten"
)

// Our game constants
const (
	windowWidth, windowHeight = 800, 600
	maxUint                   = ^uint(0)
)

//Game implements ebiten.Game interface
type Game struct {
	tick    uint
	objects []objects.Object
}

func (g *Game) Update(screen *ebiten.Image) error {
	g.tick++
	if g.tick == maxUint {
		g.tick = 0
	}

	for _, o := range g.objects {
		o.Tick(g.tick)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, o := range g.objects {
		if err := o.Draw(screen); err != nil {
			log.Fatal(err)
		}
	}
}

func (g *Game) Layout(w, h int) (int, int) {
	return w, h
}

func NewGame() *Game {
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Shooter")
	g := &Game{
		objects: []objects.Object{
			objects.NewBackground("bg_green.png"),
			objects.NewDesk("bg_wood.png"),
			objects.NewCurtains("curtain_straight.png", "curtain.png"),
		},
	}
	return g
}

func (g *Game) Run() error {
	return ebiten.RunGame(g)
}
