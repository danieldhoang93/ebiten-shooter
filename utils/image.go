package utils

import (
	"image"

	"github.com/danieldhoang93/ebiten-shooter/assets"
	"github.com/hajimehoshi/ebiten"
)

func GetImage(name string, obj *assets.Object) (*ebiten.Image, error) {
	var rectangle image.Rectangle

	for _, img := range obj.Specs.Images {
		if img.Name == name {
			rectangle = image.Rect(img.X, img.Y, img.X+img.W, img.Y+img.H)
			break
		}
	}

	img := obj.Image.SubImage(rectangle).(*ebiten.Image)
	return img, nil
}
