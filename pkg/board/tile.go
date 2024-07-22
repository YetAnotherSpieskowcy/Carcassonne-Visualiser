package board

import rl "github.com/gen2brain/raylib-go/raylib"

type Tile struct {
	Position rl.Vector2
	Size     rl.Vector2
	Color    rl.Color
}

func NewTile(pos rl.Vector2, color rl.Color) Tile {
	return Tile{
		Position: pos,
		Size:     rl.NewVector2(tileSize, tileSize),
		Color:    color,
	}
}

func (t Tile) DrawTile() {
	rl.DrawRectangleV(t.Position, t.Size, t.Color)
	rl.DrawRectangleLines(int32(t.Position.X), int32(t.Position.Y), int32(t.Size.X), int32(t.Size.Y), rl.Gray)
}
