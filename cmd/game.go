package main

import (
	"bytes"
	"fmt"
	"image"
	"log"

	_ "image/png" // to load png images

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	gauge "local.packages/gauge"
)

const (
	screenWidth  = 480
	screenHeight = 320
)

func NewGame() *Game {
	gauge := gauge.NewGauge(50, 50, 100)

	img, _, err := image.Decode(bytes.NewReader(bk_png))
	if err != nil {
		log.Println("failed to load dot png,", err)
	}
	bkImage := ebiten.NewImageFromImage(img)

	return &Game{gauge: gauge, bk: bkImage}
}

type Game struct {
	gauge   *gauge.Gauge
	counter int
	bk      *ebiten.Image
}

func (g *Game) Update() error {
	g.counter++
	g.gauge.Update(float64(g.counter))

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.bk, &ebiten.DrawImageOptions{})
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Counter: %v", g.counter))
	g.gauge.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
