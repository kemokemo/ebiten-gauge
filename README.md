# Ebiten Gauge

This is a package for displaying a simple gauge.  
It is intended to be used in games that use [ebiten](https://ebiten.org/).

// todo: sample image

## How to use

First of all, you need create an instance of the gauge.

```go
// position
x := 50
y := 80
// gauge's max value
max := 100

// create gauge instance
gauge1 := gauge.NewGauge(x, y, max)

// If you want to specify the dot's color of this gauge,
// you can use NewGaugeWithColor function
// 
gaugeGray := gauge.NewGaugeWithColor(x, y, max, color.RGBA{R: 80, G: 80, B: 80, A: 255})
```

In standard operation, the dot of gauge blinks when the current value reaches the MAX value. Blinking can be turned off for gauges that do not require blinking or are mainly used for decreasing values.

```go
gauge1 := gauge.NewGauge(x, y, max)

// turn off blinking
gauge1.SetBlink(false)
```

Second, you updates on the `ebiten.game`'s update function. Please call `Update` function with the current value for the gauge.

```go
func (g *Game) Update() error {
 // currentValue is updated by your game logic

 g.gauge1.Update(float64(currentValue))

 // some logic...
 return nil
}
```

Third, draw the gauge on the Draw function of your game.

```go
func (g *Game) Draw(screen *ebiten.Image) {
 // draw some other items

 g.gauge1.Draw(screen)

 // draw some other items
}
```

That's all!

Please check [the sample app](https://github.com/kemokemo/ebiten-gauge/tree/main/cmd) for more detail.

### Basic

## License

Apache-2.0 License

## Author

kemokemo
