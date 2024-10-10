package factory

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	engineModifier "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature/modifier"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Visualiser/pkg/board/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Visualiser/pkg/board/feature/meeple"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var cityColor = rl.NewColor(124, 67, 39, 255)

func TopCity(f elements.PlacedFeature) feature.Feature {
	hasShield := f.ModifierType == engineModifier.Shield
	cityFeature := feature.New(cityColor)

	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(0, 0), rl.NewVector2(20, 15), rl.NewVector2(20, 0)})
	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(60, 0), rl.NewVector2(40, 0), rl.NewVector2(40, 15)})
	cityFeature.AddRectangle(rl.NewVector2(20, 0), rl.NewVector2(20, 15))

	if hasShield {
		cityFeature.AddModifier(Shield(rl.NewVector2(15, 4)))
	}
	if hasMeeple(f) {
		cityFeature.AddMeeple(rl.NewVector2(30, meeple.RadiusWithMargin), f.Meeple.PlayerID)
	}

	return cityFeature
}

func BottomCity(f elements.PlacedFeature) feature.Feature {
	hasShield := f.ModifierType == engineModifier.Shield
	cityFeature := feature.New(cityColor)

	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(0, 60), rl.NewVector2(20, 60), rl.NewVector2(20, 45)})
	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(60, 60), rl.NewVector2(40, 45), rl.NewVector2(40, 60)})
	cityFeature.AddRectangle(rl.NewVector2(20, 45), rl.NewVector2(20, 15))

	if hasShield {
		cityFeature.AddModifier(Shield(rl.NewVector2(45, 3)))
	}
	if hasMeeple(f) {
		cityFeature.AddMeeple(rl.NewVector2(30, 60 - meeple.RadiusWithMargin), f.Meeple.PlayerID)
	}

	return cityFeature
}

func RightCity(f elements.PlacedFeature) feature.Feature {
	hasShield := f.ModifierType == engineModifier.Shield
	cityFeature := feature.New(cityColor)

	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(60, 0), rl.NewVector2(45, 20), rl.NewVector2(60, 20)})
	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(60, 60), rl.NewVector2(60, 40), rl.NewVector2(45, 40)})
	cityFeature.AddRectangle(rl.NewVector2(45, 20), rl.NewVector2(15, 20))

	if hasShield {
		cityFeature.AddModifier(Shield(rl.NewVector2(49, 20)))
	}
	if hasMeeple(f) {
		cityFeature.AddMeeple(rl.NewVector2(60 - meeple.RadiusWithMargin, 30), f.Meeple.PlayerID)
	}

	return cityFeature
}

func LeftCity(f elements.PlacedFeature) feature.Feature {
	hasShield := f.ModifierType == engineModifier.Shield
	cityFeature := feature.New(cityColor)

	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(0, 0), rl.NewVector2(0, 20), rl.NewVector2(15, 20)})
	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(0, 60), rl.NewVector2(15, 40), rl.NewVector2(0, 40)})
	cityFeature.AddRectangle(rl.NewVector2(0, 20), rl.NewVector2(15, 20))

	if hasShield {
		cityFeature.AddModifier(Shield(rl.NewVector2(5, 20)))
	}
	if hasMeeple(f) {
		cityFeature.AddMeeple(rl.NewVector2(meeple.RadiusWithMargin, 30), f.Meeple.PlayerID)
	}

	return cityFeature
}

func TopRightCity(f elements.PlacedFeature) feature.Feature {
	hasShield := f.ModifierType == engineModifier.Shield
	cityFeature := feature.New(cityColor)

	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(0, 0), rl.NewVector2(20, 10), rl.NewVector2(60, 0)})
	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(60, 0), rl.NewVector2(50, 40), rl.NewVector2(60, 60)})
	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(20, 10), rl.NewVector2(50, 40), rl.NewVector2(60, 0)})

	if hasShield {
		cityFeature.AddModifier(Shield(rl.NewVector2(50, 5)))
	}
	if hasMeeple(f) {
		cityFeature.AddMeeple(rl.NewVector2(60 - 18, 18), f.Meeple.PlayerID)
	}

	return cityFeature
}

func TopLeftCity(f elements.PlacedFeature) feature.Feature {
	hasShield := f.ModifierType == engineModifier.Shield
	cityFeature := feature.New(cityColor)

	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(0, 0), rl.NewVector2(40, 10), rl.NewVector2(60, 0)})
	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(0, 60), rl.NewVector2(10, 40), rl.NewVector2(0, 0)})
	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(40, 10), rl.NewVector2(0, 0), rl.NewVector2(10, 40)})

	if hasShield {
		cityFeature.AddModifier(Shield(rl.NewVector2(10, 5)))
	}
	if hasMeeple(f) {
		cityFeature.AddMeeple(rl.NewVector2(18, 18), f.Meeple.PlayerID)
	}

	return cityFeature
}

func BottomRightCity(f elements.PlacedFeature) feature.Feature {
	hasShield := f.ModifierType == engineModifier.Shield
	cityFeature := feature.New(cityColor)

	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(0, 60), rl.NewVector2(60, 60), rl.NewVector2(20, 50)})
	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(60, 0), rl.NewVector2(50, 20), rl.NewVector2(60, 60)})
	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(50, 20), rl.NewVector2(20, 50), rl.NewVector2(60, 60)})

	if hasShield {
		cityFeature.AddModifier(Shield(rl.NewVector2(50, 50)))
	}
	if hasMeeple(f) {
		cityFeature.AddMeeple(rl.NewVector2(60 - 18, 60 - 18), f.Meeple.PlayerID)
	}

	return cityFeature
}

func BottomLeftCity(f elements.PlacedFeature) feature.Feature {
	hasShield := f.ModifierType == engineModifier.Shield
	cityFeature := feature.New(cityColor)

	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(0, 0), rl.NewVector2(0, 60), rl.NewVector2(10, 20)})
	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(0, 60), rl.NewVector2(60, 60), rl.NewVector2(40, 50)})
	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(40, 50), rl.NewVector2(10, 20), rl.NewVector2(0, 60)})

	if hasShield {
		cityFeature.AddModifier(Shield(rl.NewVector2(5, 50)))
	}
	if hasMeeple(f) {
		cityFeature.AddMeeple(rl.NewVector2(18, 60 - 18), f.Meeple.PlayerID)
	}

	return cityFeature
}

func TopBottomCity(f elements.PlacedFeature) feature.Feature {
	hasShield := f.ModifierType == engineModifier.Shield
	cityFeature := feature.New(cityColor)

	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(0, 0), rl.NewVector2(15, 10), rl.NewVector2(15, 0)})
	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(60, 0), rl.NewVector2(45, 0), rl.NewVector2(45, 10)})
	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(0, 60), rl.NewVector2(15, 60), rl.NewVector2(15, 50)})
	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(45, 60), rl.NewVector2(60, 60), rl.NewVector2(45, 50)})
	cityFeature.AddRectangle(rl.NewVector2(15, 0), rl.NewVector2(30, 60))

	if hasShield {
		cityFeature.AddModifier(Shield(rl.NewVector2(15, 5)))
	}
	if hasMeeple(f) {
		cityFeature.AddMeeple(rl.NewVector2(30, 30), f.Meeple.PlayerID)
	}

	return cityFeature
}

func LeftRightCity(f elements.PlacedFeature) feature.Feature {
	hasShield := f.ModifierType == engineModifier.Shield
	cityFeature := feature.New(cityColor)

	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(0, 0), rl.NewVector2(0, 15), rl.NewVector2(10, 15)})
	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(0, 45), rl.NewVector2(0, 60), rl.NewVector2(10, 45)})
	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(60, 0), rl.NewVector2(50, 15), rl.NewVector2(60, 15)})
	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(60, 45), rl.NewVector2(50, 45), rl.NewVector2(60, 60)})
	cityFeature.AddRectangle(rl.NewVector2(0, 15), rl.NewVector2(60, 30))

	if hasShield {
		cityFeature.AddModifier(Shield(rl.NewVector2(5, 26)))
	}
	if hasMeeple(f) {
		cityFeature.AddMeeple(rl.NewVector2(30, 30), f.Meeple.PlayerID)
	}

	return cityFeature
}

func LeftTopRightCity(f elements.PlacedFeature) feature.Feature {
	hasShield := f.ModifierType == engineModifier.Shield
	cityFeature := feature.New(cityColor)

	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(0, 45), rl.NewVector2(0, 60), rl.NewVector2(10, 45)})
	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(60, 45), rl.NewVector2(50, 45), rl.NewVector2(60, 60)})
	cityFeature.AddRectangle(rl.NewVector2(0, 0), rl.NewVector2(60, 45))

	if hasShield {
		cityFeature.AddModifier(Shield(rl.NewVector2(50, 5)))
	}
	if hasMeeple(f) {
		cityFeature.AddMeeple(rl.NewVector2(30, meeple.RadiusWithMargin), f.Meeple.PlayerID)
	}

	return cityFeature
}

func TopRightBottomCity(f elements.PlacedFeature) feature.Feature {
	hasShield := f.ModifierType == engineModifier.Shield
	cityFeature := feature.New(cityColor)

	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(0, 0), rl.NewVector2(15, 10), rl.NewVector2(15, 0)})
	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(0, 60), rl.NewVector2(15, 60), rl.NewVector2(15, 50)})
	cityFeature.AddRectangle(rl.NewVector2(15, 0), rl.NewVector2(45, 60))

	if hasShield {
		cityFeature.AddModifier(Shield(rl.NewVector2(50, 5)))
	}
	if hasMeeple(f) {
		cityFeature.AddMeeple(rl.NewVector2(60-meeple.RadiusWithMargin, 30), f.Meeple.PlayerID)
	}

	return cityFeature
}

func RightBottomLeftCity(f elements.PlacedFeature) feature.Feature {
	hasShield := f.ModifierType == engineModifier.Shield
	cityFeature := feature.New(cityColor)

	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(0, 0), rl.NewVector2(0, 15), rl.NewVector2(10, 15)})
	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(60, 0), rl.NewVector2(50, 15), rl.NewVector2(60, 15)})
	cityFeature.AddRectangle(rl.NewVector2(0, 15), rl.NewVector2(60, 45))

	if hasShield {
		cityFeature.AddModifier(Shield(rl.NewVector2(50, 50)))
	}
	if hasMeeple(f) {
		cityFeature.AddMeeple(rl.NewVector2(30, 60-meeple.RadiusWithMargin), f.Meeple.PlayerID)
	}

	return cityFeature
}

func BottomLeftTopCity(f elements.PlacedFeature) feature.Feature {
	hasShield := f.ModifierType == engineModifier.Shield
	cityFeature := feature.New(cityColor)

	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(60, 0), rl.NewVector2(45, 0), rl.NewVector2(45, 10)})
	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(45, 60), rl.NewVector2(60, 60), rl.NewVector2(45, 50)})
	cityFeature.AddRectangle(rl.NewVector2(0, 0), rl.NewVector2(45, 60))

	if hasShield {
		cityFeature.AddModifier(Shield(rl.NewVector2(10, 5)))
	}
	if hasMeeple(f) {
		cityFeature.AddMeeple(rl.NewVector2(meeple.RadiusWithMargin, 30), f.Meeple.PlayerID)
	}

	return cityFeature
}
