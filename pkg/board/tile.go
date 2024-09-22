package board

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Visualiser/pkg/board/feature"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Tile struct {
	position       position.Position
	color          rl.Color
	borderColor    rl.Color
	features       []feature.Feature
	hidesMeeplesAt []position.Position
}

func NewTile(pos position.Position, color rl.Color, borderColor rl.Color) Tile {
	return Tile{
		position:    pos,
		color:       color,
		borderColor: borderColor,
		features:    make([]feature.Feature, 0),
	}
}

func (tile *Tile) AddFeature(newFeature feature.Feature) {
	tile.features = append(tile.features, newFeature)
}

func (tile Tile) calculateLocationOnBoard(offset rl.Vector2) rl.Vector2 {
	location := rl.Vector2{}
	location.X = float32((boardSize-tileSize)/2 + tile.position.X()*tileSize + int16(offset.X)*tileSize)
	location.Y = float32((boardSize-tileSize)/2 - tile.position.Y()*tileSize - int16(offset.Y)*tileSize)
	return location
}

func (tile Tile) DrawTile(offset rl.Vector2, hideMeeples bool) {
	pos := tile.calculateLocationOnBoard(offset)
	rl.DrawRectangleV(pos, rl.NewVector2(tileSize, tileSize), tile.color)

	for _, f := range tile.features {
		f.Draw(pos, hideMeeples)
	}

	rl.DrawRectangleLines(int32(pos.X), int32(pos.Y), tileSize, tileSize, tile.borderColor)
}
