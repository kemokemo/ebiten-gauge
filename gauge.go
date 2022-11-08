package gauge

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Gauge struct {
	x, y     int
	max      float64
	percent  float64
	dot      *ebiten.Image
	dotOp    []*blinkingOp
	interval int
	counter  int
	bType    BlinkingType
}

const (
	w           = 5
	h           = 10
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
	dot := ebiten.NewImage(w, h)
	dot.Fill(dotClr)

	ops := []*blinkingOp{}
	for i := 0; i < dotNum; i++ {
		bOp := newBlinkingOp(float64(firstOffset+x+(w+xInterval)*i), float64(y+yInterval))
		ops = append(ops, bOp)
	}

	return &Gauge{x: x, y: y, max: max, dotOp: ops, dot: dot, interval: 2, bType: BlinkingOnMax}
}

func (g *Gauge) SetBlinkingType(bType BlinkingType) {
	g.bType = bType
}

// Update updates the gauge appearance with the v value of arg.
func (g *Gauge) Update(v float64) {
	g.percent = v / g.max * 100
	g.counter++
	if g.counter > g.interval {
		g.counter = 0
	}

	if g.isBlinkUpdate() {
		for index := 0; index < len(g.dotOp); index++ {
			g.dotOp[index].update()
		}
	}
}

func (g *Gauge) isBlinkUpdate() bool {
	if g.bType == BlinkingOnMax {
		return g.percent >= 100 && g.counter >= g.interval
	} else if g.bType == BlinkingOnZero {
		return g.percent <= 0 && g.counter >= g.interval
	} else {
		return false
	}
}

func (g *Gauge) Draw(screen *ebiten.Image) {
	for index := 0; index < len(g.dotOp); index++ {
		if g.percent > float64((10 * index)) {
			screen.DrawImage(g.dot, g.dotOp[index].Op)
		}
	}
}
