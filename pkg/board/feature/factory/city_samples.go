package factory

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Visualiser/pkg/board/feature"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func TopCity() feature.Feature {
	cityFeature := feature.New()

	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(0, 0), rl.NewVector2(20, 10), rl.NewVector2(20, 0)}, rl.DarkBrown)
	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(60, 0), rl.NewVector2(40, 0), rl.NewVector2(40, 10)}, rl.DarkBrown)
	cityFeature.AddRectangle(rl.NewVector2(20, 0), rl.NewVector2(20, 10), rl.DarkBrown)

	return cityFeature
}

func BottomCity() feature.Feature {
	cityFeature := feature.New()

	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(0, 60), rl.NewVector2(20, 60), rl.NewVector2(20, 50)}, rl.DarkBrown)
	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(60, 60), rl.NewVector2(40, 50), rl.NewVector2(40, 60)}, rl.DarkBrown)
	cityFeature.AddRectangle(rl.NewVector2(20, 50), rl.NewVector2(20, 10), rl.DarkBrown)

	return cityFeature
}

func RightCity() feature.Feature {
	cityFeature := feature.New()

	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(60, 0), rl.NewVector2(50, 20), rl.NewVector2(60, 20)}, rl.DarkBrown)
	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(60, 60), rl.NewVector2(60, 40), rl.NewVector2(50, 40)}, rl.DarkBrown)
	cityFeature.AddRectangle(rl.NewVector2(50, 20), rl.NewVector2(10, 20), rl.DarkBrown)

	return cityFeature
}

func LeftCity() feature.Feature {
	cityFeature := feature.New()

	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(0, 0), rl.NewVector2(0, 20), rl.NewVector2(10, 20)}, rl.DarkBrown)
	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(0, 60), rl.NewVector2(10, 40), rl.NewVector2(0, 40)}, rl.DarkBrown)
	cityFeature.AddRectangle(rl.NewVector2(0, 20), rl.NewVector2(10, 20), rl.DarkBrown)

	return cityFeature
}

func TopRightCity() feature.Feature {
	cityFeature := feature.New()

	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(0, 0), rl.NewVector2(20, 10), rl.NewVector2(60, 0)}, rl.DarkBrown)
	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(60, 0), rl.NewVector2(50, 40), rl.NewVector2(60, 60)}, rl.DarkBrown)
	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(20, 10), rl.NewVector2(50, 40), rl.NewVector2(60, 0)}, rl.DarkBrown)

	return cityFeature
}

func TopLeftCity() feature.Feature {
	cityFeature := feature.New()

	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(0, 0), rl.NewVector2(40, 10), rl.NewVector2(60, 0)}, rl.DarkBrown)
	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(0, 60), rl.NewVector2(10, 40), rl.NewVector2(0, 0)}, rl.DarkBrown)
	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(40, 10), rl.NewVector2(0, 0), rl.NewVector2(10, 40)}, rl.DarkBrown)

	return cityFeature
}

func BottomRightCity() feature.Feature {
	cityFeature := feature.New()

	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(0, 60), rl.NewVector2(60, 60), rl.NewVector2(20, 50)}, rl.DarkBrown)
	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(60, 0), rl.NewVector2(50, 20), rl.NewVector2(60, 60)}, rl.DarkBrown)
	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(50, 20), rl.NewVector2(20, 50), rl.NewVector2(60, 60)}, rl.DarkBrown)

	return cityFeature
}

func BottomLeftCity() feature.Feature {
	cityFeature := feature.New()

	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(0, 0), rl.NewVector2(0, 60), rl.NewVector2(10, 20)}, rl.DarkBrown)
	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(0, 60), rl.NewVector2(60, 60), rl.NewVector2(40, 50)}, rl.DarkBrown)
	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(40, 50), rl.NewVector2(10, 20), rl.NewVector2(0, 60)}, rl.DarkBrown)

	return cityFeature
}

func TopBottomCity() feature.Feature {
	cityFeature := feature.New()

	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(0, 0), rl.NewVector2(15, 10), rl.NewVector2(15, 0)}, rl.DarkBrown)
	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(60, 0), rl.NewVector2(45, 0), rl.NewVector2(45, 10)}, rl.DarkBrown)
	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(0, 60), rl.NewVector2(15, 60), rl.NewVector2(15, 50)}, rl.DarkBrown)
	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(45, 60), rl.NewVector2(60, 60), rl.NewVector2(45, 50)}, rl.DarkBrown)
	cityFeature.AddRectangle(rl.NewVector2(15, 0), rl.NewVector2(30, 60), rl.DarkBrown)

	return cityFeature
}

func LeftRightCity() feature.Feature {
	cityFeature := feature.New()

	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(0, 0), rl.NewVector2(0, 15), rl.NewVector2(10, 15)}, rl.DarkBrown)
	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(0, 45), rl.NewVector2(0, 60), rl.NewVector2(10, 45)}, rl.DarkBrown)
	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(60, 0), rl.NewVector2(50, 15), rl.NewVector2(60, 15)}, rl.DarkBrown)
	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(60, 45), rl.NewVector2(50, 45), rl.NewVector2(60, 60)}, rl.DarkBrown)
	cityFeature.AddRectangle(rl.NewVector2(0, 15), rl.NewVector2(60, 30), rl.DarkBrown)

	return cityFeature
}

func LeftTopRightCity() feature.Feature {
	cityFeature := feature.New()

	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(0, 45), rl.NewVector2(0, 60), rl.NewVector2(10, 45)}, rl.DarkBrown)
	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(60, 45), rl.NewVector2(50, 45), rl.NewVector2(60, 60)}, rl.DarkBrown)
	cityFeature.AddRectangle(rl.NewVector2(0, 0), rl.NewVector2(60, 45), rl.DarkBrown)

	return cityFeature
}

func TopRightBottomCity() feature.Feature {
	cityFeature := feature.New()

	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(0, 0), rl.NewVector2(15, 10), rl.NewVector2(15, 0)}, rl.DarkBrown)
	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(0, 60), rl.NewVector2(15, 60), rl.NewVector2(15, 50)}, rl.DarkBrown)
	cityFeature.AddRectangle(rl.NewVector2(15, 0), rl.NewVector2(45, 60), rl.DarkBrown)

	return cityFeature
}

func RightBottomLeftCity() feature.Feature {
	cityFeature := feature.New()

	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(0, 0), rl.NewVector2(0, 15), rl.NewVector2(10, 15)}, rl.DarkBrown)
	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(60, 0), rl.NewVector2(50, 15), rl.NewVector2(60, 15)}, rl.DarkBrown)
	cityFeature.AddRectangle(rl.NewVector2(0, 15), rl.NewVector2(60, 45), rl.DarkBrown)

	return cityFeature
}

func BottomLeftTopCity() feature.Feature {
	cityFeature := feature.New()

	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(60, 0), rl.NewVector2(45, 0), rl.NewVector2(45, 10)}, rl.DarkBrown)
	cityFeature.AddTriangle([]rl.Vector2{rl.NewVector2(45, 60), rl.NewVector2(60, 60), rl.NewVector2(45, 50)}, rl.DarkBrown)
	cityFeature.AddRectangle(rl.NewVector2(0, 0), rl.NewVector2(45, 60), rl.DarkBrown)

	return cityFeature
}
