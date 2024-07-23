package board

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Tile struct {
	Position position.Position
	Color    rl.Color
}

func NewTile(pos position.Position, color rl.Color) Tile {
	return Tile{
		Position: pos,
		Color:    color,
	}
}

func (tile Tile) calculateLocationOnBoard(offset rl.Vector2) rl.Vector2 {
	location := rl.Vector2{}
	location.X = float32((boardSize-tileSize)/2 + tile.Position.X()*tileSize - int16(offset.X)*tileSize)
	location.Y = float32((boardSize-tileSize)/2 - tile.Position.Y()*tileSize + int16(offset.Y)*tileSize)
	return location
}

func (tile Tile) DrawTile(offset rl.Vector2) {
	pos := tile.calculateLocationOnBoard(offset)
	rl.DrawRectangleV(pos, rl.NewVector2(tileSize, tileSize), tile.Color)
	rl.DrawRectangleLines(int32(pos.X), int32(pos.Y), tileSize, tileSize, rl.Gray)
}
