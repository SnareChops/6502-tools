package visuals

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font/basicfont"
)

func Button(label string) *ebiten.Image {
	img, err := ebiten.NewImage(300, 150, ebiten.FilterDefault)
	// img, err := ebiten.NewImage(7*len(label), 13, ebiten.FilterDefault)
	if err != nil {
		panic(err)
	}
	text.Draw(img, label, basicfont.Face7x13, 0, 0, color.Black)
	return img
}
