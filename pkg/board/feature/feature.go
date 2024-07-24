package feature

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Visualiser/pkg/board/feature/components"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Visualiser/pkg/board/feature/modifier"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Feature struct {
	rectangles []components.Rectangle
	triangles  []components.Triangle
	color      rl.Color

	modifiers []modifier.Modifier
}

func New(color rl.Color) Feature {
	return Feature{
		rectangles: make([]components.Rectangle, 0),
		triangles:  make([]components.Triangle, 0),
		color:      color,
		modifiers:  make([]modifier.Modifier, 0),
	}
}

func (feature *Feature) AddRectangle(offsetOnTile rl.Vector2, size rl.Vector2) {
	feature.rectangles = append(feature.rectangles, components.NewRectangle(offsetOnTile, size))
}

func (feature *Feature) AddTriangle(offsetsOnTile []rl.Vector2) {
	feature.triangles = append(feature.triangles, components.NewTriangle(offsetsOnTile))
}

func (feature *Feature) AddModifier(newModifier modifier.Modifier) {
	feature.modifiers = append(feature.modifiers, newModifier)
}

func (feature Feature) Draw(tilePosition rl.Vector2) {
	for _, rectangle := range feature.rectangles {
		rectangle.Draw(tilePosition, feature.color)
	}
	for _, triangle := range feature.triangles {
		triangle.Draw(tilePosition, feature.color)
	}
	for _, modifier := range feature.modifiers {
		modifier.Draw(tilePosition)
	}
}
