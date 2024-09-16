package factory

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Visualiser/pkg/board/feature"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Monastery() feature.Feature {
	monasteryFeature := feature.New(rl.Black)
	monasteryFeature.AddRectangle(rl.NewVector2(20, 20), rl.NewVector2(20, 20))
	return monasteryFeature
}

var (
	topRoadPosition    = rl.NewVector2(25, 0)
	rightRoadPosition  = rl.NewVector2(35, 25)
	bottomRoadPosition = rl.NewVector2(25, 35)
	leftRoadPosition   = rl.NewVector2(0, 25)

	roadConnectorPosition = rl.NewVector2(25, 25)

	roadConnectorSize = rl.NewVector2(10, 10)

	verticalHalfRoadSize   = rl.NewVector2(10, 25)
	horizontalHalfRoadSize = rl.NewVector2(25, 10)
)

func Road(s side.Side) feature.Feature {
	roadFeature := feature.New(rl.DarkGray)

	edgeCtr := 0

	for _, edge := range side.PrimarySides {
		if s&edge == edge {
			edgeCtr++
			switch edge {
			case side.Top:
				roadFeature.AddRectangle(topRoadPosition, verticalHalfRoadSize)
			case side.Right:
				roadFeature.AddRectangle(rightRoadPosition, horizontalHalfRoadSize)
			case side.Bottom:
				roadFeature.AddRectangle(bottomRoadPosition, verticalHalfRoadSize)
			case side.Left:
				roadFeature.AddRectangle(leftRoadPosition, horizontalHalfRoadSize)
			}
		}
	}
	if edgeCtr > 1 {
		roadFeature.AddRectangle(roadConnectorPosition, roadConnectorSize)
	}

	return roadFeature
}

func oneEdgeCity(s side.Side, hasShield bool) feature.Feature {
	if s&side.Top == side.Top {
		return TopCity(hasShield)
	} else if s&side.Right == side.Right {
		return RightCity(hasShield)
	} else if s&side.Bottom == side.Bottom {
		return BottomCity(hasShield)
	} else {
		return LeftCity(hasShield)
	}
}

func fourEdgeCity(hasShield bool) feature.Feature {
	cityFeature := feature.New(cityColor)
	cityFeature.AddRectangle(rl.NewVector2(0, 0), rl.NewVector2(60, 60))
	if hasShield {
		cityFeature.AddModifier(Shield(rl.NewVector2(50, 5)))
	}
	return cityFeature
}

func cornerCity(s side.Side, hasShield bool) feature.Feature {
	if s&side.Top == side.Top {
		if s&side.Right == side.Right {
			return TopRightCity(hasShield)
		} else {
			return TopLeftCity(hasShield)
		}
	} else {
		if s&side.Right == side.Right {
			return BottomRightCity(hasShield)
		} else {
			return BottomLeftCity(hasShield)
		}
	}
}

func mirrorCity(s side.Side, hasShield bool) feature.Feature {
	if s&side.Top == side.Top {
		return TopBottomCity(hasShield)
	} else {
		return LeftRightCity(hasShield)
	}
}

func threeEdgeCity(s side.Side, hasShield bool) feature.Feature {
	if s&(side.Top|side.Left|side.Bottom) == (side.Top | side.Left | side.Bottom) {
		return BottomLeftTopCity(hasShield)
	} else if s&(side.Right|side.Left|side.Bottom) == (side.Right | side.Left | side.Bottom) {
		return RightBottomLeftCity(hasShield)
	} else if s&(side.Top|side.Left|side.Right) == (side.Top | side.Left | side.Right) {
		return LeftTopRightCity(hasShield)
	} else {
		return TopRightBottomCity(hasShield)
	}
}

func City(s side.Side, hasShield bool) feature.Feature {
	cityFeature := feature.New(rl.DarkBrown)
	edgesNumber := s.GetCardinalDirectionsLength()
	switch edgesNumber {
	case 1:
		return oneEdgeCity(s, hasShield)
	case 2:
		if s == s.Mirror() {
			return mirrorCity(s, hasShield)
		} else {
			return cornerCity(s, hasShield)
		}
	case 3:
		return threeEdgeCity(s, hasShield)
	case 4:
		return fourEdgeCity(hasShield)
	}
	return cityFeature
}
