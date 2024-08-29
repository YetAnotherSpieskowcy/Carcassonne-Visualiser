package components

import rl "github.com/gen2brain/raylib-go/raylib"

type Rectangle struct {
	offsetOnTile rl.Vector2
	size         rl.Vector2
}

func NewRectangle(offsetOnTile rl.Vector2, size rl.Vector2) Rectangle {
	return Rectangle{
		offsetOnTile: offsetOnTile,
		size:         size,
	}
}

func (rectangle Rectangle) Draw(tilePosition rl.Vector2, color rl.Color) {
	rl.DrawRectangleV(rl.Vector2Add(rectangle.offsetOnTile, tilePosition), rectangle.size, color)
}
