package feature

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Visualiser/pkg/board/feature/drawable"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Visualiser/pkg/board/feature/meeple"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Visualiser/pkg/board/feature/modifier"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Feature struct {
	drawable drawable.Drawable

	modifiers []modifier.Modifier

	meeples       []meeple.Meeple
}

func New(color rl.Color) Feature {
	return Feature{
		drawable:  drawable.New(color),
		modifiers: make([]modifier.Modifier, 0),
		meeples:   make([]meeple.Meeple, 0),
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

func (feature *Feature) AddMeeple(offsetOnTile rl.Vector2, playerID elements.ID) {
	var color rl.Color
	// You might ask - why care about so many player colors?
	// Because it's a huge help for testing :)
	switch playerID {
	// Carcassonne base game - 5 players
	case 1:
		color = rl.SkyBlue
	case 2:
		color = rl.Green
	case 3:
		color = rl.Yellow
	case 4:
		color = rl.Red
	case 5:
		// normally black but that would conflict with monastery color
		color = rl.White
	// Carcassonne - Big Box (2006)
	case 6:
		color = rl.LightGray
	// Carcassonne - Big Box 5 (2014)
	case 7:
		color = rl.Pink
	case 8:
		color = rl.Purple
	default:
		panic("more players than expected")
	}
	meeple := meeple.New(color)
	meeple.AddCircle(offsetOnTile)
	feature.meeples = append(feature.meeples, meeple)
}

func (feature Feature) Draw(tilePosition rl.Vector2, hideMeeples bool) {
	feature.drawable.Draw(tilePosition)
	for _, modifier := range feature.modifiers {
		modifier.Draw(tilePosition)
	}
	if !hideMeeples {
		for _, meeple := range feature.meeples {
			meeple.Draw(tilePosition)
		}
	}
}
