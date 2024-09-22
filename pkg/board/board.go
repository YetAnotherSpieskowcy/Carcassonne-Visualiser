package board

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
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

	maxTileX int16
	minTileX int16
	maxTileY int16
	minTileY int16

	nextMoves []Tile
	tiles     []Tile

	hiddenMeeples map[position.Position]struct{}
}

func NewBoard(startTile Tile) Board {
	nextMoves := make([]Tile, 0)

	tiles := make([]Tile, 0)
	tiles = append(tiles, startTile)

	return Board{
		screenWidth:   boardSize,
		screenHeight:  boardSize,
		minRange:      rl.NewVector2(-6, -6),
		maxRange:      rl.NewVector2(6, 6),
		offset:        rl.NewVector2(0, 0),
		nextMoves:     nextMoves,
		tiles:         tiles,
		hiddenMeeples: map[position.Position]struct{}{},
	}
}

func (board *Board) MoveBoard(direction rl.Vector2) {
	newOffset := rl.Vector2Add(board.offset, direction)

	if float32(board.maxTileX)+newOffset.X >= board.minRange.X && float32(board.minTileX)+newOffset.X <= board.maxRange.X &&
		float32(board.maxTileY)+newOffset.Y >= board.minRange.Y && float32(board.minTileY)+newOffset.Y <= board.maxRange.Y {

		board.offset = newOffset
	}
}

func (board *Board) NextMove(nextTile Tile, nextTilePlaced bool) bool {
	defer board.findTileExtremes()

	if len(board.nextMoves) > 0 {
		tileIndex := len(board.nextMoves) - 1
		board.tiles = append(board.tiles, board.nextMoves[tileIndex])
		board.nextMoves = board.nextMoves[:tileIndex]
		board.applyHiddenMeeplesFromCurrentTile()
		return false
	} else if !nextTilePlaced {
		board.tiles = append(board.tiles, nextTile)
		board.applyHiddenMeeplesFromCurrentTile()
		return true
	}
	return true
}

func (board *Board) UpdateHidesMeepleAtForCurrentTile(scoreReport elements.ScoreReport) {
	currentTile := board.tiles[len(board.tiles) - 1]

	for _, meeples := range scoreReport.ReturnedMeeples {
		for _, meeple := range meeples {
			currentTile.hidesMeeplesAt = append(currentTile.hidesMeeplesAt, meeple.Position)
		}
	}

	board.tiles[len(board.tiles) - 1] = currentTile
	board.applyHiddenMeeplesFromCurrentTile()
}

func (board *Board) applyHiddenMeeplesFromCurrentTile() {
	currentTile := board.tiles[len(board.tiles) - 1]

	for _, pos := range currentTile.hidesMeeplesAt {
		board.hiddenMeeples[pos] = struct{}{}
	}
}

func (board *Board) revertHiddenMeeplesFromCurrentTile() {
	currentTile := board.tiles[len(board.tiles) - 1]

	for _, pos := range currentTile.hidesMeeplesAt {
		delete(board.hiddenMeeples, pos)
	}
}

func (board *Board) PreviousMove() {
	if len(board.tiles) > 1 { // we leave starting tile
		tileIndex := len(board.tiles) - 1
		board.nextMoves = append(board.nextMoves, board.tiles[tileIndex])
		board.revertHiddenMeeplesFromCurrentTile()
		board.tiles = board.tiles[:tileIndex]
		board.findTileExtremes()
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
		if tile.position.X()+int16(board.offset.X) >= int16(board.minRange.X) && tile.position.X()+int16(board.offset.X) <= int16(board.maxRange.X) &&
			tile.position.Y()+int16(board.offset.Y) >= int16(board.minRange.Y) && tile.position.Y()+int16(board.offset.Y) <= int16(board.maxRange.Y) {
			_, hideMeeples := board.hiddenMeeples[tile.position]
			tile.DrawTile(board.offset, hideMeeples)
		}
	}

	rl.EndDrawing()
}

func (board *Board) findTileExtremes() {
	board.maxTileX, board.minTileX = 0, 0
	board.maxTileY, board.minTileY = 0, 0

	for _, tile := range board.tiles {
		if tile.position.X() > board.maxTileX {
			board.maxTileX = tile.position.X()
		} else if tile.position.X() < board.minTileX {
			board.minTileX = tile.position.X()
		}

		if tile.position.Y() > board.maxTileY {
			board.maxTileY = tile.position.Y()
		} else if tile.position.Y() < board.minTileY {
			board.minTileY = tile.position.Y()
		}
	}

	if float32(board.maxTileX)+board.offset.X < board.minRange.X {
		board.offset.X = board.minRange.X - float32(board.maxTileX)
	} else if float32(board.minTileX)+board.offset.X > board.maxRange.X {
		board.offset.X = board.maxRange.X - float32(board.minTileX)
	}

	if float32(board.maxTileY)+board.offset.Y < board.minRange.Y {
		board.offset.Y = board.minRange.Y - float32(board.maxTileY)
	} else if float32(board.minTileY)+board.offset.Y > board.maxRange.Y {
		board.offset.Y = board.maxRange.Y - float32(board.minTileY)
	}
}
