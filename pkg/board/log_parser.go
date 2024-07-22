package board

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func calculateLocationOnBoard(pos position.Position) rl.Vector2 {
	location := rl.Vector2{}
	location.X = float32((boardSize-tileSize)/2 + pos.X()*tileSize)
	location.Y = float32((boardSize-tileSize)/2 - pos.Y()*tileSize)
	return location
}

func ParseStartEntry() Tile {
	return NewTile(calculateLocationOnBoard(position.New(0, 0)), rl.Green)
}

func ParsePlaceTileEntry(position position.Position, color rl.Color) Tile {
	return NewTile(calculateLocationOnBoard(position), color)
}
