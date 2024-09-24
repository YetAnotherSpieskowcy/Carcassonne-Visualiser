package components

import rl "github.com/gen2brain/raylib-go/raylib"

type Circle struct {
	offsetOnTile rl.Vector2
	radius       float32
}

func NewCircle(offsetOnTile rl.Vector2, radius float32) Circle {
	return Circle{
		offsetOnTile: offsetOnTile,
		radius:       radius,
	}
}

func (circle Circle) Draw(tilePosition rl.Vector2, color rl.Color) {
	rl.DrawCircleV(rl.Vector2Add(circle.offsetOnTile, tilePosition), circle.radius, color)
}
