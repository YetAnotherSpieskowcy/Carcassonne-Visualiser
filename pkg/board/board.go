package board

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	tileSize  = 60
	boardSize = 800
	offset    = boardSize % tileSize
)

type Board struct {
	screenWidth  int32
	screenHeight int32

	nextMoves []Tile
	tiles     []Tile
}

func NewBoard() Board {
	nextMoves := make([]Tile, 0)
	nextMoves = append(nextMoves, ParsePlaceTileEntry(position.New(0, 1), rl.Blue))
	nextMoves = append(nextMoves, ParsePlaceTileEntry(position.New(-6, 6), rl.Red))

	tiles := make([]Tile, 0)
	tiles = append(tiles, ParseStartEntry())

	return Board{
		screenWidth:  boardSize,
		screenHeight: boardSize,
		nextMoves:    nextMoves,
		tiles:        tiles,
	}
}

func (board *Board) AddTile(tile Tile) {
	board.tiles = append(board.tiles, tile)
}

func (board *Board) NextMove() {
	if len(board.nextMoves) > 0 {
		tileIndex := len(board.nextMoves) - 1
		board.tiles = append(board.tiles, board.nextMoves[tileIndex])
		board.nextMoves = board.nextMoves[:tileIndex]
	}
}

func (board *Board) PreviousMove() {
	if len(board.tiles) > 1 { // we leave starting tile
		tileIndex := len(board.tiles) - 1
		board.nextMoves = append(board.nextMoves, board.tiles[tileIndex])
		board.tiles = board.tiles[:tileIndex]
	}
}

func (board Board) Draw() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.RayWhite)
	// Draw grid lines
	for i := int32(0); i < board.screenWidth/tileSize+1; i++ {
		rl.DrawLineV(
			rl.NewVector2(float32(tileSize*i)+offset/2, offset/2),
			rl.NewVector2(float32(tileSize*i)+offset/2, float32(board.screenHeight)-offset/2),
			rl.LightGray,
		)
	}

	for i := int32(0); i < board.screenHeight/tileSize+1; i++ {
		rl.DrawLineV(
			rl.NewVector2(offset/2, float32(tileSize*i)+offset/2),
			rl.NewVector2(float32(board.screenWidth)-offset/2, float32(tileSize*i)+offset/2),
			rl.LightGray,
		)
	}

	// Draw tiles
	for _, tile := range board.tiles {
		tile.DrawTile()
	}

	rl.EndDrawing()
}
