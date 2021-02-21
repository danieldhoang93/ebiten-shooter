package objects

import (
	"log"
	"math"

	"github.com/danieldhoang93/ebiten-shooter/assets"
	"github.com/danieldhoang93/ebiten-shooter/utils"
	"github.com/hajimehoshi/ebiten"
)

const (
	duckName        = "duck_outline_target_white.png"
	ducksXSpeed     = 3.5
	ducksYSpeed     = 1
	ducksMaxOffsetY = 16 //max vertical movement for animation
)

type duck struct {
	img            *ebiten.Image
	offsetY        float64
	offsetX        float64
	initialOffsetY float64
	yDirection     direction //whether duck is moving up or down
	onScreen       bool
}

//newDuck generates a new duck with an initial vertical position
func newDuck(initialOffsetY int) *duck {
	img, err := utils.GetImage(duckName, assets.Objects)

	if err != nil {
		log.Fatalf("drawing %s: %v", duckName, err)
	}

	w, _ := img.Size()

	return &duck{
		img:            img,
		offsetY:        0,
		offsetX:        float64(-w),
		initialOffsetY: float64(initialOffsetY),
		yDirection:     down,
		onScreen:       true,
	}
}

func (d *duck) Tick(screen *ebiten.Image, _ uint) {
	//horizontal movement
	d.offsetX = d.offsetX + ducksXSpeed

	//when the duck is off the screen, flip bool
	screenWidth, _ := screen.Size()
	if d.offsetX > float64(screenWidth) {
		d.onScreen = false
	}

	//calculate the vertical direction and offset for the animation
	if ducksMaxOffsetY-math.Abs(d.offsetY) < 0 {
		d.yDirection = d.yDirection.invert()
	}
	d.offsetY = d.offsetY + float64(d.yDirection)*ducksYSpeed
}

func (d *duck) Draw(image *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(d.offsetX, d.offsetY+d.initialOffsetY)
	image.DrawImage(d.img, op)

	return nil
}

func (d *duck) OnScreen() bool {
	return d.onScreen
}
