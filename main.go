package main

import (
	"ebiten_mouse_move_effect/utils"
	"ebiten_mouse_move_effect/vec2d"
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

// Global variables
var (
	Black         color.Color = color.RGBA{R: 0, G: 0, B: 0, A: 255}
	Lime          color.Color = color.RGBA{R: 0, G: 255, B: 0, A: 255}
	ScreenWidth               = 800
	ScreenHeight              = 600
	Mouse         vec2d.Vec2D
	PreviousMouse vec2d.Vec2D
)

func (game *Game) Update() error {
	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	screen.Fill(Black)

	if utils.MouseHasMoved(&Mouse, &PreviousMouse) {
		ebitenutil.DrawCircle(screen, Mouse.X, Mouse.Y, 30, Lime)
	}
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func main() {
	rand.Seed(time.Now().UnixNano())

	game := &Game{}

	ebiten.SetWindowPosition(600, 200)
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Mouse Move Effects")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
