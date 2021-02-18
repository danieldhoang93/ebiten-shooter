package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Our game constants
const (
	screenWidth, screenHeight = 800, 600
	maxUint                   = ^uint(0)
)

//Game implements ebiten.Game interface
type Game struct {
	tick uint
}

// Create our empty variables
var (
	err        error
	background *ebiten.Image
	spaceShip  *ebiten.Image
	playerOne  player
)

// Create the player class
type player struct {
	image      *ebiten.Image
	xPos, yPos float64
	speed      float64
}

// Run this code once at startup
func init() {
	background, _, err = ebitenutil.NewImageFromFile("../assets/space.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	spaceShip, _, err = ebitenutil.NewImageFromFile("../assets/spaceship.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	playerOne = player{spaceShip, screenWidth / 2.0, screenHeight / 2.0, 5}
}

// Move the player depending on which key is pressed
func movePlayer() {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		playerOne.yPos -= playerOne.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		playerOne.yPos += playerOne.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		playerOne.xPos -= playerOne.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		playerOne.xPos += playerOne.speed
	}
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update(screen *ebiten.Image) error {
	movePlayer()
	return nil
}

//Draw draws the game screen and is called every frame
func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(background, op)

	playerOp := &ebiten.DrawImageOptions{}
	playerOp.GeoM.Translate(playerOne.xPos, playerOne.yPos)
	screen.DrawImage(playerOne.image, playerOp)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWdith, screenHeight int) {
	return 600, 480
}

func main() {
	game := &Game{}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

}
