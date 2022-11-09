package gauge

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func newBlinkingOp(x, y float64) *blinkingOp {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	op.ColorM.Scale(1.0, 1.0, 1.0, 1.0)
	return &blinkingOp{Op: op, increasing: false, alpha: 1.0, counter: 10}
}

type blinkingOp struct {
	Op *ebiten.DrawImageOptions

	alpha      float64
	counter    int
	increasing bool
}

func (b *blinkingOp) update() {
	if b.increasing && b.counter < 10 {
		b.counter++
	} else if b.increasing && b.counter == 10 {
		b.increasing = false
		b.counter--
	} else if !b.increasing && 0 < b.counter {
		b.counter--
	} else {
		b.increasing = true
		b.counter++
	}
	b.alpha = 0.1 * float64(b.counter)

	b.Op.ColorM.Reset()
	b.Op.ColorM.Scale(1.0, 1.0, 1.0, b.alpha)
}

func (b *blinkingOp) clear() {
	b.alpha = 1.0
	b.counter = 10
	b.increasing = false
	b.Op.ColorM.Reset()
	b.Op.ColorM.Scale(1.0, 1.0, 1.0, b.alpha)
}
