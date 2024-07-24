package pkg

import (
	"fmt"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/logger"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Visualiser/pkg/addons"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Visualiser/pkg/board"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	board        board.Board
	controlsInfo addons.Info

	logs           <-chan logger.Entry
	nextTile       board.Tile
	nextTilePlaced bool
}

func (game *Game) Init(filename string) {
	fileLogger, _ := logger.NewFromFile(filename)

	game.logs = fileLogger.ReadLogs()

	game.board = board.NewBoard(board.ParseStartEntry(<-game.logs))
	game.nextTile = board.ParsePlaceTileEntry(<-game.logs)
	game.nextTilePlaced = false

	game.controlsInfo = addons.NewInfo(
		"A - Previous move, D - Next move\nArrows - Move board",
		rl.NewVector2(10, 815),
	)

}

func (game *Game) Update(nextMove bool) {
	if nextMove {
		readNewEntry := game.board.NextMove(game.nextTile, game.nextTilePlaced)
		if readNewEntry {
			game.nextTilePlaced = true
			val, ok := <-game.logs
			if ok {
				fmt.Println("reading new tile")
				game.nextTile = board.ParsePlaceTileEntry(val)
				game.nextTilePlaced = false
			}
		}
	} else {
		game.board.PreviousMove()
	}
}

func (game *Game) MoveBoard(direction rl.Vector2) {
	game.board.MoveBoard(direction)
}

func (game *Game) Draw() {
	game.board.Draw()
	game.controlsInfo.DrawInfo()
}
