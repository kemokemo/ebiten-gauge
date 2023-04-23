package gauge

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const defaultInterval = 10

func NewBlinkingOp(dotClr color.Color) *BlinkingOp {
	op := &ebiten.DrawImageOptions{}
	op.ColorScale.ScaleWithColor(dotClr)
	return &BlinkingOp{Op: op, increasing: false, alpha: 1.0, counter: defaultInterval, interval: defaultInterval}
}

// DrawImageOption extension to make an element blink.
//
// Set the blinking speed with SetInterval and update the blinking state with Update
// while drawing the element using the property Op.
type BlinkingOp struct {
	Op *ebiten.DrawImageOptions

	alpha      float32
	counter    int
	interval   int
	increasing bool
}

// SetInterval sets the interval of blinking.
// Larger values will cause it to blink more slowly.
func (b *BlinkingOp) SetInterval(interval int) {
	b.interval = interval
}

// Update updates the internal values to blink.
func (b *BlinkingOp) Update() {
	if b.increasing && b.counter < b.interval {
		b.counter++
	} else if b.increasing && b.counter == b.interval {
		b.increasing = false
		b.counter--
	} else if !b.increasing && 0 < b.counter {
		b.counter--
	} else {
		b.increasing = true
		b.counter++
	}
	b.alpha = 0.1 * float32(b.counter)

	b.Op.ColorScale.SetA(b.alpha)
}

// Clear clears internal values except the interval value.
// If you want to change the blinking interval, please use SetInterval method.
func (b *BlinkingOp) Clear() {
	b.alpha = 1.0
	b.counter = defaultInterval
	b.increasing = false
	b.Op.ColorScale.SetA(b.alpha)
}
