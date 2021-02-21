package objects

import (
	"fmt"
	"math"

	"github.com/danieldhoang93/ebiten-shooter/assets"
	"github.com/danieldhoang93/ebiten-shooter/utils"
	"github.com/hajimehoshi/ebiten"
)

type background struct {
	name string
}

func NewBackground(img string) Object {
	return &background{
		name: img,
	}
}

func (b *background) Tick(image *ebiten.Image, tick uint) {}

func (b *background) Draw(target *ebiten.Image) error {
	img, err := utils.GetImage(b.name, assets.Stall)
	if err != nil {
		return fmt.Errorf("drawing %s: %v", b.name, err)
	}

	//repeat image to fit background
	trgtW, trgtH := target.Size()
	bgW, bgH := img.Size()

	x := int(math.Ceil(float64(trgtW) / float64(bgW)))
	y := int(math.Ceil(float64(trgtH) / float64(bgH)))

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			op := &ebiten.DrawImageOptions{}
			tx := i * bgW
			ty := j * bgH
			op.GeoM.Translate(float64(tx), float64(ty))
			target.DrawImage(img, op)
		}
	}

	return nil
}

func (b *background) OnScreen() bool {
	return true
}
