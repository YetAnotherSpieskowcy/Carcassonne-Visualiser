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

	logs            LogReader
	nextEntry       logger.Entry
	nextEntryExists bool
	moveCtr         uint32
	moveCtrMax      uint32
}

func (game *Game) Init(filename string) {
	logger, err := logger.NewLogReader(filename)
	if err != nil {
		panic(err)
	}
	game.logs = logger

	startEntry, _ := game.logs.ReadEntry()
	startTile, numOfPlayer := ParseStartEntry(startEntry)

	game.scoreInfo = addons.NewScoreInfo(numOfPlayer, rl.NewVector2(810, 10))
	game.board = board.NewBoard(startTile)

	game.nextEntry, game.nextEntryExists = game.logs.ReadEntry()

	game.moveCtr = 0
	game.moveCtrMax = 0

	game.controlsInfo = addons.NewInfo(
		"A - Previous move, D - Next move\nArrows - Move board",
		rl.NewVector2(10, 815),
	)
}

func (game *Game) Update(nextMove bool) {
	if nextMove {
		game.nextMove()
	} else {
		game.previousMove()
	}
}

func (game *Game) nextMove() {
	moveWasMade := false

	if game.moveCtr < game.moveCtrMax {
		// replay a move from previous moves history (already loaded from logs)
		game.board.NextMove()
		moveWasMade = true

	} else {
		// play a new move
		if !game.nextEntryExists {
			game.nextEntry, game.nextEntryExists = game.logs.ReadEntry()
		}

		shouldProcessNextEntry := true
		for shouldProcessNextEntry && game.nextEntryExists {
			if game.nextEntry.Event == logger.PlaceTileEvent {
				moveWasMade = true
			}
			game.processEntry(game.nextEntry)
			game.nextEntry, game.nextEntryExists = game.logs.ReadEntry()

			// we want to process all score events immediately after a tile was placed, but we don't want to place two tiles in a row
			if game.nextEntry.Event == logger.PlaceTileEvent {
				shouldProcessNextEntry = false
			}
		}
	}

	if moveWasMade {
		game.incrementMoveCtr()
	}
	game.scoreInfo.NextScores(game.moveCtr)
}

func (game *Game) processEntry(entry logger.Entry) {
	if entry.Event == logger.PlaceTileEvent {
		tile := ParsePlaceTileEntry(entry)
		game.board.NextNewMove(tile)

	} else if entry.Event == logger.ScoreEvent {
		scoreReport := ParseScoreEntry(entry)
		for _, meeples := range scoreReport.ReturnedMeeples {
			for _, meeple := range meeples {
				game.board.ResetTile(meeple.Position)
			}
		}
		game.scoreInfo.UpdateScores(scoreReport, game.moveCtr+1)
	} else {
		panic("only PlaceTileEvent and ScoreEvent entries are currently supported")
	}
}

func (game *Game) previousMove() {
	if game.moveCtr == 0 {
		return
	}

	game.board.PreviousMove()
	game.scoreInfo.PreviousScores(game.moveCtr)
	game.decrementMoveCtr()
}

func (game *Game) incrementMoveCtr() {
	game.moveCtr++
	if game.moveCtr > game.moveCtrMax {
		game.moveCtrMax = game.moveCtr
	}
}

func (game *Game) decrementMoveCtr() {
	game.moveCtr--
}

func (game *Game) MoveBoard(direction rl.Vector2) {
	game.board.MoveBoard(direction)
}

func (game *Game) Draw() {
	game.board.Draw()
	game.controlsInfo.Show()
	game.scoreInfo.Show(game.moveCtr)
}
