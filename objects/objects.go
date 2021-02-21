package objects

import "github.com/hajimehoshi/ebiten"

type Object interface {
	Tick(*ebiten.Image, uint) // tell the object a new tick happened
	Draw(*ebiten.Image) error // draw the object
	OnScreen() bool           //false when object is off screen
}

//custom type for direction
type direction int

const (
	right direction = 1
	left  direction = -1
	up    direction = -1
	down  direction = 1
)

func (d direction) invert() direction {
	return -d
}
