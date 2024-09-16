package factory

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	CityColor      = rl.NewColor(124, 67, 39, 255)
	RoadColor      = rl.DarkGray
	FieldColor     = rl.DarkGreen
	MonasteryColor = rl.Black
)

var MeepleColors = []rl.Color{
	rl.White, // playerID = 0, should never appear
	rl.Blue,
	rl.Red,
}
