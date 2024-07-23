package board

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func ParseStartEntry() Tile {
	return NewTile(position.New(0, 0), rl.Green)
}

func ParsePlaceTileEntry(position position.Position, color rl.Color) Tile {
	return NewTile(position, color)
}
