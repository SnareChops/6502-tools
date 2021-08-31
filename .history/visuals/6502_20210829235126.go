package visuals

import (
	"image"
	"image/color"

	"github.com/SnareChops/6502-tools/sim"
	"github.com/SnareChops/6502-tools/utils"
	"github.com/hajimehoshi/ebiten"
)

var Image6502Options = &ebiten.DrawImageOptions{}

func init() {
	Image6502Options.GeoM.Translate(426, 0)
}

func Image6502(model *sim.Model) *ebiten.Image {
	// Create new Image
	img := image.NewRGBA(image.Rect(0, 0, 428, 200))
	// Draw rectangle for processor body
	utils.Rect(img, image.Rect(0, 75, 427, 175), color.Black)
	// Draw pin rectangles
	for y := 65; y < 176; y += 110 {
		for x := 20; x < 401; x += 20 {
			utils.Rect(img, image.Rect(x, y, x+10, y+10), color.Black)
		}
	}
	// Convert to ebiten image
	result, err := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		panic(err)
	}
	// return ebiten image
	return result
}

type VisualModel struct {
	*sim.Model
}

func (m *VisualModel) DrawPins(img *image.RGBA) {
	for i := 1; i < 21; i++ {
		y := 65
		x := (i * 20) + 20
		utils.Rect(img, image.Rect(x, y, x+10, y+10), color.Black)
	}
}
