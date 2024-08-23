package modifier

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Visualiser/pkg/board/feature/drawable"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Modifier struct {
	drawable drawable.Drawable
	name     string
}

func New(color rl.Color, name string) Modifier {
	return Modifier{
		drawable: drawable.New(color),
	}
}

func (modifier *Modifier) AddRectangle(offset rl.Vector2, size rl.Vector2) {
	modifier.drawable.AddRectangle(offset, size)
}

func (modifier *Modifier) AddTriangle(offsetsOnTile []rl.Vector2) {
	modifier.drawable.AddTriangle(offsetsOnTile)
}

func (modifier Modifier) Draw(tilePosition rl.Vector2) {
	modifier.drawable.Draw(tilePosition)
}
