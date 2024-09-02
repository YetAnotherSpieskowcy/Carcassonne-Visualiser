package drawable

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Visualiser/pkg/board/feature/drawable/components"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Drawable struct {
	rectangles     []components.Rectangle
	triangles      []components.Triangle
	color          rl.Color
	customColor    rl.Color
	useCustomColor bool
}

func New(color rl.Color) Drawable {
	return Drawable{
		rectangles:     make([]components.Rectangle, 0),
		triangles:      make([]components.Triangle, 0),
		color:          color,
		customColor:    rl.Black,
		useCustomColor: false,
	}
}

func (drawable *Drawable) SetCustomColor(color rl.Color) {
	drawable.customColor = color
	drawable.useCustomColor = true
}

func (drawable *Drawable) ClearCustomColor() {
	drawable.useCustomColor = false
}

func (drawable *Drawable) AddRectangle(offsetOnTile rl.Vector2, size rl.Vector2) {
	drawable.rectangles = append(drawable.rectangles, components.NewRectangle(offsetOnTile, size))
}

func (drawable *Drawable) AddTriangle(offsetsOnTile []rl.Vector2) {
	drawable.triangles = append(drawable.triangles, components.NewTriangle(offsetsOnTile))
}

func (drawable Drawable) Draw(tilePosition rl.Vector2) {
	var color rl.Color
	if drawable.useCustomColor {
		color = drawable.customColor
	} else {
		color = drawable.color
	}

	for _, rectangle := range drawable.rectangles {
		rectangle.Draw(tilePosition, color)
	}
	for _, triangle := range drawable.triangles {
		triangle.Draw(tilePosition, color)
	}
}
