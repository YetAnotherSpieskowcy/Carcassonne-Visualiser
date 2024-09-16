package feature

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Visualiser/pkg/board/feature/drawable"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Visualiser/pkg/board/feature/modifier"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Feature struct {
	drawable drawable.Drawable

	modifiers []modifier.Modifier
}

func New(color rl.Color) Feature {
	return Feature{
		drawable:  drawable.New(color),
		modifiers: make([]modifier.Modifier, 0),
	}
}

func (feature *Feature) AddRectangle(offsetOnTile rl.Vector2, size rl.Vector2) {
	feature.drawable.AddRectangle(offsetOnTile, size)
}

func (feature *Feature) AddTriangle(offsetsOnTile []rl.Vector2) {
	feature.drawable.AddTriangle(offsetsOnTile)
}

func (feature *Feature) AddModifier(newModifier modifier.Modifier) {
	feature.modifiers = append(feature.modifiers, newModifier)
}

func (feature Feature) Draw(tilePosition rl.Vector2) {
	feature.drawable.Draw(tilePosition)
	for _, modifier := range feature.modifiers {
		modifier.Draw(tilePosition)
	}
}
