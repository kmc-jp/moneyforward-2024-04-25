package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// 画面の大きさ(画素数)を定義しておく
var (
	worldWidth  = 640
	worldHeight = 480
)

var (
	Magnification    float32 = 1.03 // 拡大率
	ColorAttenuation float32 = 0.98 // 色の減衰率
)

type Bubble struct {
	R float32
	C float32 // Gray scale color
}

type Game struct {
	b *Bubble
}

func (g *Game) Update() error {
	g.b.R *= Magnification
	g.b.C = g.b.C * ColorAttenuation
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	vector.StrokeCircle(screen, 100, 100, g.b.R, 1, color.Gray{Y: uint8(g.b.C)}, true)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return worldWidth, worldHeight
}

func main() {

	b := &Bubble{R: 10, C: 255}

	ebiten.SetTPS(20)

	ebiten.SetWindowSize(worldWidth, worldHeight)
	ebiten.SetWindowTitle("Ebitengine Bubble Art")
	if err := ebiten.RunGame(&Game{b: b}); err != nil {
		log.Fatal(err)
	}
}
