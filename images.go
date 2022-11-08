package gauge

import (
	"bytes"
	"image"
	"log"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	dotImage *ebiten.Image
)

func init() {
	img, _, err := image.Decode(bytes.NewReader(dot_png))
	if err != nil {
		log.Println("failed to load dot png,", err)
	}
	dotImage = ebiten.NewImageFromImage(img)
}
