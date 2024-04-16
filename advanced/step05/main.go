package main

import (
	"container/list"
	"fmt"
	"image/color"
	"log"
	"math"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// 画面の大きさ(画素数)を定義しておく
var (
	worldWidth  = 640
	worldHeight = 480
)

var (
	MaxBubbleNum     int     = 100
	Magnification    float32 = 1.03 // 拡大率
	ColorAttenuation float32 = 0.98 // 色の減衰率
)

type Bubble struct {
	X float32
	Y float32
	R float32
	C float32 // Gray scale color
}

func NewBubble(centerX, centerY int) *Bubble {
	// 外円の半径
	R := worldHeight / 2

	// 中心からの距離を乱数で決める
	a := float64(rand.IntN(1000) + 1) // 1-1000
	d := float32(math.Log10(a)/3) * float32(R)
	x := d
	y := float32(0)

	// 回転
	theta := rand.Float32() * 2 * math.Pi
	x2 := x*float32(math.Cos(float64(theta))) - y*float32(math.Sin(float64(theta)))
	y2 := x*float32(math.Sin(float64(theta))) + y*float32(math.Cos(float64(theta)))

	// 中心に平行移動
	x2 += float32(centerX)
	y2 += float32(centerY)

	b := &Bubble{X: float32(x2), Y: float32(y2), R: 10, C: 255}
	return b
}

type Game struct {
	bubbles *list.List
}

func (g *Game) Update() error {
	if g.bubbles.Len() >= MaxBubbleNum {
		g.bubbles.Remove(g.bubbles.Front())
	}

	for e := g.bubbles.Front(); e != nil; e = e.Next() {
		b := e.Value.(*Bubble)
		b.R *= Magnification
		b.C = b.C * ColorAttenuation
	}

	g.bubbles.PushBack(NewBubble(worldWidth/2, worldHeight/2))
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for e := g.bubbles.Front(); e != nil; e = e.Next() {
		b := e.Value.(*Bubble)
		vector.StrokeCircle(screen, b.X, b.Y, b.R, 1, color.Gray{Y: uint8(b.C)}, true)
	}
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Bubble Count: %d", g.bubbles.Len()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return worldWidth, worldHeight
}

func main() {

	ebiten.SetTPS(20)

	ebiten.SetWindowSize(worldWidth, worldHeight)
	ebiten.SetWindowTitle("Ebitengine Bubble Art")
	if err := ebiten.RunGame(&Game{bubbles: list.New()}); err != nil {
		log.Fatal(err)
	}
}
