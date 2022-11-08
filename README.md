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
gauge := gauge.NewGauge(x, y, max)

// If you want to specify the dot's color of this gauge,
// you can use NewGaugeWithColor function
// 
gaugeGray := gauge.NewGaugeWithColor(x, y, max, color.RGBA{R: 80, G: 80, B: 80, A: 255})
```

Second, you updates on the `ebiten.game`'s update function. Please call `Update` function with the current value for the gauge.

```go
func (g *Game) Update() error {
 // currentValue is updated by your game logic

 g.gauge.Update(float64(currentValue))

 // some logic...
 return nil
}
```

That's all! If the current value reaches the `max` value, gauge's dots will blink.

Please check [the sample app](https://github.com/kemokemo/ebiten-gauge/tree/main/cmd) for more detail.

### Basic

## License

Apache-2.0 License

## Author

kemokemo
