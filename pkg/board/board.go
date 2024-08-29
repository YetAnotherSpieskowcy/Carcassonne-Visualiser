package board

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	tileSize    = 60
	boardSize   = 800
	boardOffset = boardSize % tileSize
)

type Board struct {
	screenWidth  int32
	screenHeight int32
	offset       rl.Vector2
	minRange     rl.Vector2
	maxRange     rl.Vector2

	nextMoves []Tile
	tiles     []Tile
}

func NewBoard(startTile Tile) Board {
	nextMoves := make([]Tile, 0)

	tiles := make([]Tile, 0)
	tiles = append(tiles, startTile)

	return Board{
		screenWidth:  boardSize,
		screenHeight: boardSize,
		minRange:     rl.NewVector2(-6, -6),
		maxRange:     rl.NewVector2(6, 6),
		offset:       rl.NewVector2(0, 0),
		nextMoves:    nextMoves,
		tiles:        tiles,
	}
}

func (board *Board) MoveBoard(direction rl.Vector2) {
	board.offset = rl.Vector2Add(board.offset, direction)
	board.minRange = rl.Vector2Subtract(board.minRange, direction)
	board.maxRange = rl.Vector2Subtract(board.maxRange, direction)
}

func (board *Board) NextMove(nextTile Tile, nextTilePlaced bool) bool {
	if len(board.nextMoves) > 0 {
		tileIndex := len(board.nextMoves) - 1
		board.tiles = append(board.tiles, board.nextMoves[tileIndex])
		board.nextMoves = board.nextMoves[:tileIndex]
		return false
	} else if !nextTilePlaced {
		board.tiles = append(board.tiles, nextTile)
		return true
	}
	return true
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
			rl.NewVector2(float32(tileSize*i)+boardOffset/2, boardOffset/2),
			rl.NewVector2(float32(tileSize*i)+boardOffset/2, float32(board.screenHeight)-boardOffset/2),
			rl.LightGray,
		)
	}

	for i := int32(0); i < board.screenHeight/tileSize+1; i++ {
		rl.DrawLineV(
			rl.NewVector2(boardOffset/2, float32(tileSize*i)+boardOffset/2),
			rl.NewVector2(float32(board.screenWidth)-boardOffset/2, float32(tileSize*i)+boardOffset/2),
			rl.LightGray,
		)
	}

	// Draw tiles
	for _, tile := range board.tiles {
		// Draw tile only if it is visible
		if tile.position.X() >= int16(board.minRange.X) && tile.position.X() <= int16(board.maxRange.X) &&
			tile.position.Y() >= int16(board.minRange.Y) && tile.position.Y() <= int16(board.maxRange.Y) {
			tile.DrawTile(board.offset)
		}
	}

	rl.EndDrawing()
}
