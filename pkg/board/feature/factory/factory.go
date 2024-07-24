package factory

import (
	"fmt"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Visualiser/pkg/board/feature"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Monastery() feature.Feature {
	monasteryFeature := feature.New()
	monasteryFeature.AddRectangle(rl.NewVector2(20, 20), rl.NewVector2(20, 20), rl.Black)
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
	roadFeature := feature.New()

	edgeCtr := 0

	for _, edge := range side.PrimarySides {
		if s&edge == edge {
			edgeCtr++
			switch edge {
			case side.Top:
				roadFeature.AddRectangle(topRoadPosition, verticalHalfRoadSize, rl.DarkGray)
			case side.Right:
				roadFeature.AddRectangle(rightRoadPosition, horizontalHalfRoadSize, rl.DarkGray)
			case side.Bottom:
				roadFeature.AddRectangle(bottomRoadPosition, verticalHalfRoadSize, rl.DarkGray)
			case side.Left:
				roadFeature.AddRectangle(leftRoadPosition, horizontalHalfRoadSize, rl.DarkGray)
			}
		}
	}
	if edgeCtr > 1 {
		roadFeature.AddRectangle(roadConnectorPosition, roadConnectorSize, rl.DarkGray)
	}

	return roadFeature
}

func oneEdgeCity(s side.Side) feature.Feature {
	if s&side.Top == side.Top {
		return TopCity()
	} else if s&side.Right == side.Right {
		return RightCity()
	} else if s&side.Bottom == side.Bottom {
		return BottomCity()
	} else {
		return LeftCity()
	}
}

func fourEdgeCity() feature.Feature {
	cityFeature := feature.New()
	cityFeature.AddRectangle(rl.NewVector2(0, 0), rl.NewVector2(60, 60), rl.DarkBrown)
	return cityFeature
}

var ()

func cornerCity(s side.Side) feature.Feature {
	if s&side.Top == side.Top {
		if s&side.Right == side.Right {
			fmt.Println("TopRight")
			return TopRightCity()
		} else {
			fmt.Println("TopLeft")
			return TopLeftCity()
		}
	} else {
		if s&side.Right == side.Right {
			fmt.Println("BottomRight")
			return BottomRightCity()
		} else {
			fmt.Println("BottomLeft")
			return BottomLeftCity()
		}
	}
}

func mirrorCity(s side.Side) feature.Feature {
	if s&side.Top == side.Top {
		return TopBottomCity()
	} else {
		return LeftRightCity()
	}
}

func threeEdgeCity(s side.Side) feature.Feature {
	if s&(side.Top|side.Left|side.Bottom) == (side.Top | side.Left | side.Bottom) {
		return BottomLeftTopCity()
	} else if s&(side.Right|side.Left|side.Bottom) == (side.Right | side.Left | side.Bottom) {
		return RightBottomLeftCity()
	} else if s&(side.Top|side.Left|side.Right) == (side.Top | side.Left | side.Right) {
		return LeftTopRightCity()
	} else {
		return TopRightBottomCity()
	}
}

func City(s side.Side) feature.Feature {
	cityFeature := feature.New()
	edgesNumber := s.GetCardinalDirectionsLength()
	switch edgesNumber {
	case 1:
		return oneEdgeCity(s)
	case 2:
		if s == s.Mirror() {
			return mirrorCity(s)
		} else {
			return cornerCity(s)
		}
	case 3:
		return threeEdgeCity(s)
	case 4:
		return fourEdgeCity()
	}
	return cityFeature
}
