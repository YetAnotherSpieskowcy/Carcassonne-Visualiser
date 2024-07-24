package modifier

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Visualiser/pkg/board/feature/components"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Modifier struct {
	rectangles []components.Rectangle
	triangles  []components.Triangle
	color      rl.Color
}

func New(color rl.Color) Modifier {
	return Modifier{
		rectangles: make([]components.Rectangle, 0),
		triangles:  make([]components.Triangle, 0),
		color:      color,
	}
}

func (modifier *Modifier) AddRectangle(offset rl.Vector2, size rl.Vector2) {
	modifier.rectangles = append(modifier.rectangles, components.NewRectangle(offset, size))
}

func (modifier *Modifier) AddTriangle(offsetsOnTile []rl.Vector2) {
	modifier.triangles = append(modifier.triangles, components.NewTriangle(offsetsOnTile))
}

func (modifier Modifier) Draw(tilePosition rl.Vector2) {
	for _, rectangle := range modifier.rectangles {
		rectangle.Draw(tilePosition, modifier.color)
	}
	for _, triangle := range modifier.triangles {
		triangle.Draw(tilePosition, modifier.color)
	}
}
