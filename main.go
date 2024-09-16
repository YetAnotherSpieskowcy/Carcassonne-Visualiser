package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Visualiser/pkg"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 1200
	screenHeight = 875
)

func main() {
	log_filename := os.Args[1]

	if _, err := os.Stat(log_filename); errors.Is(err, os.ErrNotExist) {
		message := "File " + log_filename + " does not exist."
		fmt.Println(message)
		os.Exit(2)
	}

	game := pkg.Game{}
	game.Init(log_filename)

	rl.InitWindow(screenWidth, screenHeight, "Carcassonne-Visualiser")
	rl.SetTargetFPS(60)

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
			game.MoveBoard(rl.NewVector2(0, 1))
		case rl.KeyDown:
			game.MoveBoard(rl.NewVector2(0, -1))
		}
		game.Draw()
	}

	rl.CloseWindow()

}
