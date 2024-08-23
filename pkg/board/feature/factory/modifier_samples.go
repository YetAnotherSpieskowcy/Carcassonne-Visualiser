package factory

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Visualiser/pkg/board/feature/modifier"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Shield(offset rl.Vector2) modifier.Modifier {
	shield := modifier.New(rl.LightGray, "Shield")

	shield.AddRectangle(offset, rl.NewVector2(6, 5))
	shield.AddTriangle([]rl.Vector2{rl.NewVector2(offset.X, offset.Y+5), rl.NewVector2(offset.X+3, offset.Y+8), rl.NewVector2(offset.X+6, offset.Y+5)})

	return shield
}
