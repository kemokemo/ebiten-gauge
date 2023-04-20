package main

import (
	"bytes"
	"fmt"
	"image"
	"log"

	"image/color"
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
	gaugeMax := gauge.NewGauge(50, 50, 100)
	gaugeMax.SetBlinkInterval(5)
	gaugeMaxSize2 := gauge.NewGaugeWithScale(250, 50, 100, color.RGBA{R: 29, G: 87, B: 199, A: 255}, 2.0)
	gaugeZero := gauge.NewGaugeWithColor(50, 113, 100, color.RGBA{R: 80, G: 80, B: 80, A: 255})
	gaugeZero.SetBlink(false)

	img, _, err := image.Decode(bytes.NewReader(bk_png))
	if err != nil {
		log.Println("failed to load dot png,", err)
	}
	bkImage := ebiten.NewImageFromImage(img)

	return &Game{gaugeMax: gaugeMax, gaugeMaxSize2: gaugeMaxSize2, gaugeZero: gaugeZero, bk: bkImage, counterForZero: 100, increasing: true, decreasing: true}
}

type Game struct {
	gaugeMax       *gauge.Gauge
	gaugeMaxSize2  *gauge.Gauge
	increasing     bool
	decreasing     bool
	gaugeZero      *gauge.Gauge
	counter        int
	counterForZero int
	bk             *ebiten.Image
}

func (g *Game) Update() error {
	if g.counter > 150 && g.increasing {
		g.increasing = false
	} else if g.counter < -50 && !g.increasing {
		g.increasing = true
	} else {
		// keep flag
	}

	if g.counterForZero < -50 && g.decreasing {
		g.decreasing = false
	} else if g.counterForZero > 150 && !g.decreasing {
		g.decreasing = true
	} else {
		// keep flag
	}

	if g.increasing {
		g.counter++
	} else {
		g.counter--
	}
	g.gaugeMax.Update(float64(g.counter))
	g.gaugeMaxSize2.Update(float64(g.counter))

	if g.decreasing {
		g.counterForZero--
	} else {
		g.counterForZero++
	}
	g.gaugeZero.Update(float64(g.counterForZero))

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.bk, &ebiten.DrawImageOptions{})
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Counter: %v", g.counter))
	g.gaugeMax.Draw(screen)
	g.gaugeMaxSize2.Draw(screen)
	g.gaugeZero.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
