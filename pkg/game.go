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
	scoreInfo    addons.ScoreInfo

	logs           <-chan logger.Entry
	nextTile       board.Tile
	nextTilePlaced bool
	moveCtr        uint32
}

func (game *Game) Init(filename string) {
	fileLogger, _ := logger.NewFromFile(filename)

	game.logs = fileLogger.ReadLogs()

	startTile, numOfPlayer := ParseStartEntry(<-game.logs)

	game.scoreInfo = addons.NewScoreInfo(numOfPlayer, rl.NewVector2(810, 10))
	game.board = board.NewBoard(startTile)
	game.nextTile = ParsePlaceTileEntry(<-game.logs)
	game.nextTilePlaced = false
	game.moveCtr = 0

	game.controlsInfo = addons.NewInfo(
		"A - Previous move, D - Next move\nArrows - Move board",
		rl.NewVector2(10, 815),
	)

}

func (game *Game) Update(nextMove bool) {
	if nextMove {
		game.moveCtr++
		nextTileWasPlaced := game.nextTilePlaced
		readNewEntry := game.board.NextMove(game.nextTile, nextTileWasPlaced)
		if readNewEntry {
			game.nextTilePlaced = true
			for game.nextTilePlaced {
				entry, ok := <-game.logs
				if !ok {
					// channel already closed - no further tiles expected
					if nextTileWasPlaced {
						// `Board.NextMove()` may ask for next tile after placing one
						// to have another one ready but, if the channel is closed,
						// there won't be another move ever
						game.moveCtr--
					}
					return
				}
				if entry.Event == logger.PlaceTileEvent {
					game.nextTile = ParsePlaceTileEntry(entry)
					game.nextTilePlaced = false
				} else if entry.Event == logger.ScoreEvent {
					scoreReport := ParseScoreEntry(entry)
					game.scoreInfo.UpdateScores(scoreReport, game.moveCtr)
					game.board.UpdateHidesMeepleAtForCurrentTile(scoreReport)
				}
			}
		} else {
			// move was already in cache, the scores just need to be updated
			game.scoreInfo.NextScores(game.moveCtr)
		}
	} else if game.moveCtr > 0 {
		game.board.PreviousMove()
		game.scoreInfo.PreviousScores(game.moveCtr)
		game.moveCtr--
	}
}

func (game *Game) MoveBoard(direction rl.Vector2) {
	game.board.MoveBoard(direction)
}

func (game *Game) Draw() {
	game.board.Draw()
	game.controlsInfo.Show()
	game.scoreInfo.Show(game.moveCtr)
}
