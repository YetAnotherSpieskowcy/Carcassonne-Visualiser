package components

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Triangle struct {
	offsetsOnTile []rl.Vector2
}

/*
INFO: rl.DrawTriangle requires that the coordinates of triangle corners
were given in counterclockwise direction, therefore when passing list of
offsets of corners at tile remember to make sure that they keep a good
direction, otherwise triangle won't be drawn.
*/
func NewTriangle(offsetsOnTile []rl.Vector2) Triangle {
	return Triangle{
		offsetsOnTile: offsetsOnTile,
	}
}

func (triangle Triangle) Draw(tilePosition rl.Vector2, color rl.Color) {
	rl.DrawTriangle(rl.Vector2Add(triangle.offsetsOnTile[0], tilePosition),
		rl.Vector2Add(triangle.offsetsOnTile[1], tilePosition),
		rl.Vector2Add(triangle.offsetsOnTile[2], tilePosition),
		color)
}
