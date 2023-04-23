package gauge

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed assets/dot-filled.png
	dot_filled []byte
	//go:embed assets/dot-empty.png
	dot_empty []byte

	dotFilled *ebiten.Image
	dotEmpty  *ebiten.Image
)

func init() {
	var err error
	dotFilled, err = loadSingleImage(dot_filled)
	if err != nil {
		log.Println("failed to load dotFilled image, ", err)
		return
	}

	dotEmpty, err = loadSingleImage(dot_empty)
	if err != nil {
		log.Println("failed to load dotEmpty image, ", err)
		return
	}
}

func loadSingleImage(b []byte) (*ebiten.Image, error) {
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	return ebiten.NewImageFromImage(img), nil
}
