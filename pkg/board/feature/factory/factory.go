package factory

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	engineModifier "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature/modifier"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Visualiser/pkg/board/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Visualiser/pkg/board/feature/meeple"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func hasMeeple(f elements.PlacedFeature) bool {
	return f.Meeple.Type != elements.NoneMeeple
}

func hasShield(f elements.PlacedFeature) bool {
	return f.ModifierType == engineModifier.Shield
}

func Monastery(f elements.PlacedFeature) feature.Feature {
	monasteryFeature := feature.New(rl.Black)
	monasteryFeature.AddRectangle(rl.NewVector2(20, 20), rl.NewVector2(20, 20))
	if hasMeeple(f) {
		monasteryFeature.AddMeeple(rl.NewVector2(30, 30), f.Meeple.PlayerID)
	}
	return monasteryFeature
}

var (
	topRoadPosition    = rl.NewVector2(25, 0)
	rightRoadPosition  = rl.NewVector2(35, 25)
	bottomRoadPosition = rl.NewVector2(25, 35)
	leftRoadPosition   = rl.NewVector2(0, 25)

	topMeeplePosition    = rl.NewVector2(30, meeple.RadiusWithMargin)
	rightMeeplePosition  = rl.NewVector2(60 - meeple.RadiusWithMargin, 30)
	bottomMeeplePosition = rl.NewVector2(30, 60 - meeple.RadiusWithMargin)
	leftMeeplePosition   = rl.NewVector2(meeple.RadiusWithMargin, 30)

	roadConnectorPosition = rl.NewVector2(25, 25)

	roadConnectorSize = rl.NewVector2(10, 10)

	verticalHalfRoadSize   = rl.NewVector2(10, 25)
	horizontalHalfRoadSize = rl.NewVector2(25, 10)
)

func Road(f elements.PlacedFeature) feature.Feature {
	roadFeature := feature.New(rl.DarkGray)

	edgeCtr := 0

	s := f.Sides
	meeplePlaced := !hasMeeple(f)
	for _, edge := range side.PrimarySides {
		if s&edge == edge {
			edgeCtr++
			switch edge {
			case side.Top:
				roadFeature.AddRectangle(topRoadPosition, verticalHalfRoadSize)
				if !meeplePlaced {
					roadFeature.AddMeeple(topMeeplePosition, f.Meeple.PlayerID)
					meeplePlaced = true
				}
			case side.Right:
				roadFeature.AddRectangle(rightRoadPosition, horizontalHalfRoadSize)
				if !meeplePlaced {
					roadFeature.AddMeeple(rightMeeplePosition, f.Meeple.PlayerID)
					meeplePlaced = true
				}
			case side.Bottom:
				roadFeature.AddRectangle(bottomRoadPosition, verticalHalfRoadSize)
				if !meeplePlaced {
					roadFeature.AddMeeple(bottomMeeplePosition, f.Meeple.PlayerID)
					meeplePlaced = true
				}
			case side.Left:
				roadFeature.AddRectangle(leftRoadPosition, horizontalHalfRoadSize)
				if !meeplePlaced {
					roadFeature.AddMeeple(leftMeeplePosition, f.Meeple.PlayerID)
					meeplePlaced = true
				}
			}
		}
	}
	if edgeCtr > 1 {
		roadFeature.AddRectangle(roadConnectorPosition, roadConnectorSize)
	}

	return roadFeature
}

func oneEdgeCity(f elements.PlacedFeature) feature.Feature {
	s := f.Sides
	if s&side.Top == side.Top {
		return TopCity(f)
	} else if s&side.Right == side.Right {
		return RightCity(f)
	} else if s&side.Bottom == side.Bottom {
		return BottomCity(f)
	} else {
		return LeftCity(f)
	}
}

func fourEdgeCity(f elements.PlacedFeature) feature.Feature {
	cityFeature := feature.New(cityColor)
	cityFeature.AddRectangle(rl.NewVector2(0, 0), rl.NewVector2(60, 60))
	if hasShield(f) {
		cityFeature.AddModifier(Shield(rl.NewVector2(50, 5)))
	}
	if hasMeeple(f) {
		cityFeature.AddMeeple(rl.NewVector2(30, 30), f.Meeple.PlayerID)
	}
	return cityFeature
}

func cornerCity(f elements.PlacedFeature) feature.Feature {
	s := f.Sides

	if s&side.Top == side.Top {
		if s&side.Right == side.Right {
			return TopRightCity(f)
		} else {
			return TopLeftCity(f)
		}
	} else {
		if s&side.Right == side.Right {
			return BottomRightCity(f)
		} else {
			return BottomLeftCity(f)
		}
	}
}

func mirrorCity(f elements.PlacedFeature) feature.Feature {
	s := f.Sides

	if s&side.Top == side.Top {
		return TopBottomCity(f)
	} else {
		return LeftRightCity(f)
	}
}

func threeEdgeCity(f elements.PlacedFeature) feature.Feature {
	s := f.Sides

	if s&(side.Top|side.Left|side.Bottom) == (side.Top | side.Left | side.Bottom) {
		return BottomLeftTopCity(f)
	} else if s&(side.Right|side.Left|side.Bottom) == (side.Right | side.Left | side.Bottom) {
		return RightBottomLeftCity(f)
	} else if s&(side.Top|side.Left|side.Right) == (side.Top | side.Left | side.Right) {
		return LeftTopRightCity(f)
	} else {
		return TopRightBottomCity(f)
	}
}

func City(f elements.PlacedFeature) feature.Feature {
	s := f.Sides

	cityFeature := feature.New(rl.DarkBrown)
	edgesNumber := s.GetCardinalDirectionsLength()
	switch edgesNumber {
	case 1:
		return oneEdgeCity(f)
	case 2:
		if s == s.Mirror() {
			return mirrorCity(f)
		} else {
			return cornerCity(f)
		}
	case 3:
		return threeEdgeCity(f)
	case 4:
		return fourEdgeCity(f)
	}
	return cityFeature
}

func Field(f elements.PlacedFeature) feature.Feature {
	fieldFeature := feature.New(rl.DarkGreen)
	if !hasMeeple(f) {
		return fieldFeature
	}

	s := f.Sides

	for _, edge := range side.PrimarySides {
		if s&edge != edge {
			continue
		}

		switch edge {
		case side.Top:
			fieldFeature.AddMeeple(topMeeplePosition, f.Meeple.PlayerID)
		case side.Right:
			fieldFeature.AddMeeple(rightMeeplePosition, f.Meeple.PlayerID)
		case side.Bottom:
			fieldFeature.AddMeeple(bottomMeeplePosition, f.Meeple.PlayerID)
		case side.Left:
			fieldFeature.AddMeeple(leftMeeplePosition, f.Meeple.PlayerID)
		}
		return fieldFeature
	}

	for _, edge := range side.EdgeSides {
		if s&edge != edge {
			continue
		}

		switch edge {
		case side.TopLeftEdge:
			fieldFeature.AddMeeple(rl.NewVector2(16, 8), f.Meeple.PlayerID)
		case side.TopRightEdge:
			fieldFeature.AddMeeple(rl.NewVector2(60 - 16, 8), f.Meeple.PlayerID)

		case side.BottomLeftEdge:
			fieldFeature.AddMeeple(rl.NewVector2(16, 60 - 8), f.Meeple.PlayerID)
		case side.BottomRightEdge:
			fieldFeature.AddMeeple(rl.NewVector2(60 - 16, 60 - 8), f.Meeple.PlayerID)

		case side.LeftTopEdge:
			fieldFeature.AddMeeple(rl.NewVector2(8, 16), f.Meeple.PlayerID)
		case side.LeftBottomEdge:
			fieldFeature.AddMeeple(rl.NewVector2(8, 60 - 16), f.Meeple.PlayerID)

		case side.RightTopEdge:
			fieldFeature.AddMeeple(rl.NewVector2(60 - 8, 16), f.Meeple.PlayerID)
		case side.RightBottomEdge:
			fieldFeature.AddMeeple(rl.NewVector2(60 - 8, 60 - 16), f.Meeple.PlayerID)
		}
		return fieldFeature
	}

	return fieldFeature
}
