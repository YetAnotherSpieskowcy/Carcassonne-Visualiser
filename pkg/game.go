package pkg

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/logger"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Visualiser/pkg/addons"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Visualiser/pkg/board"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	board        board.Board
	controlsInfo addons.Info

	logger logger.FileLogger

	ctr uint8
}

func (game *Game) Init() {
	game.board = board.NewBoard()
	game.controlsInfo = addons.NewInfo(
		"D - Next move, A - Previous move",
		rl.NewVector2(225, 810),
	)
	game.ctr = 0

}

func (game *Game) Update(nextMove bool) {
	if nextMove {
		game.board.NextMove()
	} else {
		game.board.PreviousMove()
	}
}

func (game *Game) Draw() {
	game.board.Draw()
	game.controlsInfo.DrawInfo()
}
