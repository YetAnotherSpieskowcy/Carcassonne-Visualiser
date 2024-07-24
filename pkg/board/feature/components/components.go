package components

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Rectangle struct {
	offsetOnTile rl.Vector2
	size         rl.Vector2
	color        rl.Color
}

func NewRectangle(offsetOnTile rl.Vector2, size rl.Vector2, color rl.Color) Rectangle {
	return Rectangle{
		offsetOnTile: offsetOnTile,
		size:         size,
		color:        color,
	}
}

func (rectangle Rectangle) Draw(tilePosition rl.Vector2) {
	rl.DrawRectangleV(rl.Vector2Add(rectangle.offsetOnTile, tilePosition), rectangle.size, rectangle.color)
}

type Triangle struct {
	offsetsOnTile []rl.Vector2
	color         rl.Color
}

/*
INFO: rl.DrawTriangle requires that the coordinates of triangle corners
were given in counterclockwise direction, therefore when passing list of
offsets of corners at tile remember to make sure that they keep a good
direction, otherwise triangle won't be drawn.
*/
func NewTriangle(offsetsOnTile []rl.Vector2, color rl.Color) Triangle {
	return Triangle{
		offsetsOnTile: offsetsOnTile,
		color:         color,
	}
}

func (triangle Triangle) Draw(tilePosition rl.Vector2) {
	rl.DrawTriangle(rl.Vector2Add(triangle.offsetsOnTile[0], tilePosition),
		rl.Vector2Add(triangle.offsetsOnTile[1], tilePosition),
		rl.Vector2Add(triangle.offsetsOnTile[2], tilePosition),
		triangle.color)
}
