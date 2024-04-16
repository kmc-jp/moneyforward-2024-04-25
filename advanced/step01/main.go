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

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	vector.StrokeCircle(screen, 100, 100, 10, 1, color.White, true)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return worldWidth, worldHeight
}

func main() {
	ebiten.SetWindowSize(worldWidth, worldHeight)
	ebiten.SetWindowTitle("Ebitengine Bubble Art")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
