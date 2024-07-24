package feature

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Visualiser/pkg/board/feature/components"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Feature struct {
	rectangles []components.Rectangle
	triangles  []components.Triangle
}

func New() Feature {
	return Feature{
		rectangles: make([]components.Rectangle, 0),
		triangles:  make([]components.Triangle, 0),
	}
}

func (feature *Feature) AddRectangle(offsetOnTile rl.Vector2, size rl.Vector2, color rl.Color) {
	feature.rectangles = append(feature.rectangles, components.NewRectangle(offsetOnTile, size, color))
}

func (feature *Feature) AddTriangle(offsetsOnTile []rl.Vector2, color rl.Color) {
	feature.triangles = append(feature.triangles, components.NewTriangle(offsetsOnTile, color))
}

func (feature *Feature) Rotate(rotations uint8) {

}

func (feature Feature) Draw(tilePosition rl.Vector2) {
	for _, rectangle := range feature.rectangles {
		rectangle.Draw(tilePosition)
	}
	for _, triangle := range feature.triangles {
		triangle.Draw(tilePosition)
	}
}
