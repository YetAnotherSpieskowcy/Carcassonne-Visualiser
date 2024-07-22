package main

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Visualiser/pkg"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 800
	screenHeight = 850
)

func main() {
	game := pkg.Game{}
	game.Init()

	rl.InitWindow(screenWidth, screenHeight, "tests")

	for !rl.WindowShouldClose() {
		key := rl.GetKeyPressed()
		if key == rl.KeyD {
			game.Update(true)
		} else if key == rl.KeyA {
			game.Update(false)
		}
		game.Draw()
	}

	rl.CloseWindow()
}
