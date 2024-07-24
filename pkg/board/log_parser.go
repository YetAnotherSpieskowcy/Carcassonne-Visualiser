package board

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/logger"
	engineFeature "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Visualiser/pkg/board/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Visualiser/pkg/board/feature/factory"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func parseFeatures(f elements.PlacedFeature) feature.Feature {
	if f.FeatureType == engineFeature.Monastery {
		return factory.Monastery()
	} else if f.FeatureType == engineFeature.Road {
		return factory.Road(f.Sides)
	} else {
		return factory.City(f.Sides)
	}
}

func ParseStartEntry(entry logger.Entry) Tile {
	startContent := logger.ParseStartEntryContent(entry.Content)

	placedStartTile := elements.ToPlacedTile(startContent.StartingTile)

	tile := NewTile(position.New(0, 0), rl.Green)

	for _, f := range placedStartTile.Features {
		if f.FeatureType != engineFeature.Field {
			tile.AddFeature(parseFeatures(f))
		}
	}

	return tile
}

func ParsePlaceTileEntry(entry logger.Entry) Tile {
	placedTileContent := logger.ParsePlaceTileEntryContent(entry.Content)

	tile := NewTile(placedTileContent.Move.Position, rl.DarkGreen)

	for _, f := range placedTileContent.Move.Features {
		if f.FeatureType != engineFeature.Field {
			tile.AddFeature(parseFeatures(f))
		}
	}

	return tile
}
