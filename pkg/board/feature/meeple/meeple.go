package meeple

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Visualiser/pkg/board/feature/drawable"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	Radius           = 5.0
	RadiusWithMargin = Radius + 2
)

type Meeple struct {
	drawable drawable.Drawable
}

func New(color rl.Color) Meeple {
	return Meeple{
		drawable: drawable.New(color),
	}
}

func (meeple *Meeple) AddCircle(offsetOnTile rl.Vector2) {
	meeple.drawable.AddCircle(offsetOnTile, Radius)
}

func (meeple Meeple) Draw(tilePosition rl.Vector2) {
	meeple.drawable.Draw(tilePosition)
}
