package gauge

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Gauge struct {
	x, y          int
	max           float64
	percent       float64
	prevPercent   float64
	dotFilled     *ebiten.Image
	dotEmpty      *ebiten.Image
	dotOp         []*BlinkingOp
	interval      int
	counter       int
	blink         bool
	blinkFinished bool
}

const (
	w           = 5
	h           = 11
	dotNum      = 10
	xInterval   = 3
	yInterval   = 5
	firstOffset = 5
)

var (
	defaultDotColor = color.RGBA{R: 29, G: 87, B: 199, A: 255}
)

// NewGauge generates a new gauge ui component.
//
//	x,y: position
//	max: gauge's max value
func NewGauge(x, y int, max float64) *Gauge {
	return NewGaugeWithColor(x, y, max, defaultDotColor)
}

// NewGauge generates a new gauge ui component with specified colors.
//
//	x,y: position
//	max: gauge's max value
//	dotClr: dot's color
//	bkClr: background color
func NewGaugeWithColor(x, y int, max float64, dotClr color.Color) *Gauge {
	return NewGaugeWithScale(x, y, max, dotClr, 1.0)
}

func NewGaugeWithScale(x, y int, max float64, dotClr color.Color, scale float64) *Gauge {
	imgW := int(w * scale)

	ops := []*BlinkingOp{}
	for i := 0; i < dotNum; i++ {
		bOp := NewBlinkingOp(dotClr)
		bOp.Op.GeoM.Scale(scale, scale)
		bOp.Op.GeoM.Translate(float64(firstOffset+x+(imgW+xInterval)*i), float64(y+yInterval))
		ops = append(ops, bOp)
	}

	return &Gauge{x: x, y: y, max: max, dotOp: ops, dotFilled: dotFilled, dotEmpty: dotEmpty, interval: 2, blink: true}
}

func (g *Gauge) SetBlink(blink bool) {
	g.blink = blink
}

func (g *Gauge) SetBlinkInterval(interval int) {
	for index := 0; index < len(g.dotOp); index++ {
		g.dotOp[index].SetInterval(interval)
	}
}

// Update updates the gauge appearance with the v value of arg.
func (g *Gauge) Update(v float64) {
	g.prevPercent = g.percent
	g.percent = v / g.max * 100

	if g.prevPercent >= 100 && g.percent < 100 {
		g.blinkFinished = true
	} else {
		g.blinkFinished = false
	}

	g.counter++
	if g.counter > g.interval {
		g.counter = 0
	}

	if g.blink {
		g.blinkUpdate()
	}
}

func (g *Gauge) blinkUpdate() {
	if g.percent >= 100 && g.counter >= g.interval {
		for index := 0; index < len(g.dotOp); index++ {
			g.dotOp[index].Update()
		}
	}

	if g.blinkFinished {
		for index := 0; index < len(g.dotOp); index++ {
			g.dotOp[index].Clear()
		}
	}
}

func (g *Gauge) Draw(screen *ebiten.Image) {
	for index := 0; index < len(g.dotOp); index++ {
		if g.percent > float64((10 * index)) {
			screen.DrawImage(g.dotFilled, g.dotOp[index].Op)
		} else {
			screen.DrawImage(g.dotEmpty, g.dotOp[index].Op)
		}
	}
}
