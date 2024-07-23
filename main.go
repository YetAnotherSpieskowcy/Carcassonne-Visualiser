package main

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Visualiser/pkg"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 800
	screenHeight = 875
)

func main() {
	game := pkg.Game{}
	game.Init()

	rl.InitWindow(screenWidth, screenHeight, "tests")

	for !rl.WindowShouldClose() {
		key := rl.GetKeyPressed()
		switch key {
		case rl.KeyD:
			game.Update(true)
		case rl.KeyA:
			game.Update(false)
		case rl.KeyLeft:
			game.MoveBoard(rl.NewVector2(-1, 0))
		case rl.KeyRight:
			game.MoveBoard(rl.NewVector2(1, 0))
		case rl.KeyUp:
			game.MoveBoard(rl.NewVector2(0, -1))
		case rl.KeyDown:
			game.MoveBoard(rl.NewVector2(0, 1))
		}
		game.Draw()
	}

	rl.CloseWindow()
}
