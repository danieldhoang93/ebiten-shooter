package objects

import (
	"log"
	"math"
	"math/rand"

	"github.com/danieldhoang93/ebiten-shooter/assets"
	"github.com/danieldhoang93/ebiten-shooter/utils"
	"github.com/hajimehoshi/ebiten"
)

type level struct {
	img        *ebiten.Image //water waves
	imgW       int
	imgH       int
	offsetX    float64 //horizontal offset used to animate
	offsetY    float64 //vertical offset used to animate
	xDirection direction
	yDirection direction
	ducks      []*duck //current number of ducks on the screen
	maxDucks   int
}

const (
	levelXSpeed     = 1
	levelYSpeed     = .85
	levelMaxOffsetX = 100 //max horizontal movement
	levelMaxOffsetY = 16  //max vertical movement
)

func NewLevel(imgName string, maxDucks int) Object {
	img, err := utils.GetImage(imgName, assets.Stall)
	if err != nil {
		log.Fatalf("cannot get image %s: %v", imgName, err)
	}

	w, h := img.Size()

	return &level{
		img:        img,
		imgW:       w,
		imgH:       h,
		xDirection: right,
		yDirection: down,
		maxDucks:   maxDucks,
	}
}

func (level *level) Tick(image *ebiten.Image, tick uint) {
	// if the current number of ducks is below the expected number, maybe generate one
	if len(level.ducks) < level.maxDucks {
		// every half  second there's 30% possibilities to generate a duck
		if tick%10 == 0 && rand.Float64() < 0.3 {
			level.ducks = append(level.ducks, newDuck(level.imgH+65))
		}
	}

	// Update the tick of the ducks and check if they're still
	// on screen, removing from the list if not
	// Note: as we're playing with a slice while looping over
	// it, we use an external n counter and at the end of
	// the loop we reduce the slice to the final length
	// https://github.com/golang/go/wiki/SliceTricks#filter-in-place
	n := 0
	for _, d := range level.ducks {
		d.Tick(image, tick)
		if d.onScreen {
			level.ducks[n] = d
			n++
		}
	}
	level.ducks = level.ducks[:n]

	// Calculate the horizontal offset of the image.
	// First the direction:
	if level.offsetX >= levelMaxOffsetX {
		level.xDirection = level.xDirection.invert()
	} else if level.offsetX <= 0 {
		level.xDirection = right
	}
	// Then the actual calculation
	level.offsetX = level.offsetX + float64(level.xDirection)*levelXSpeed

	// Same for vertical animation
	if level.offsetY >= levelMaxOffsetY {
		level.yDirection = up
	} else if level.offsetY <= 0 {
		level.yDirection = down
	}
	level.offsetY = level.offsetY + float64(level.yDirection)*levelYSpeed
}

func (level *level) Draw(image *ebiten.Image) error {
	// Draw the ducks before the water because they must be below it
	for _, d := range level.ducks {
		d.Draw(image)
	}

	imageW, imageH := image.Size()
	// x is the number of images to draw horizontally to fill in the whole screen
	x := int(math.Ceil(float64(imageW) / float64(level.imgW)))
	// the loop starts at -1 to add an additional element
	// out of the screen on the left, that will become visible
	// during the horizontal movement
	for i := -1; i < x; i++ {
		op := &ebiten.DrawImageOptions{}
		// horizontal offset of the image, we're using multiple images to fill in the screen
		tx := i * level.imgW
		// vertically we move the image at the bottom of the screen
		ty := imageH - level.imgH
		op.GeoM.Translate(float64(tx), float64(ty))
		// apply offset to animate the image
		op.GeoM.Translate(level.offsetX, level.offsetY)

		image.DrawImage(level.img, op)
	}

	return nil
}

func (l *level) OnScreen() bool {
	return true
}
