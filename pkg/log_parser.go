package pkg

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/logger"
	engineFeature "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	engineModifier "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature/modifier"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Visualiser/pkg/board"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Visualiser/pkg/board/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Visualiser/pkg/board/feature/factory"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func parseFeatures(f elements.PlacedFeature) feature.Feature {
	var newFeature feature.Feature

	if f.FeatureType == engineFeature.Monastery {
		newFeature = factory.Monastery()
	} else if f.FeatureType == engineFeature.Road {
		newFeature = factory.Road(f.Sides)
	} else if f.FeatureType == engineFeature.Field {
		newFeature = factory.Field(f.Sides)
	} else if f.FeatureType == engineFeature.City {
		newFeature = factory.City(f.Sides, f.ModifierType == engineModifier.Shield)
	} else {
		panic("unrecognised feature type")
	}

	if f.Meeple.Type != elements.NoneMeeple {
		newFeature.SetCustomColor(factory.MeepleColors[f.Meeple.PlayerID])
	}

	return newFeature
}

func ParseStartEntry(entry logger.Entry) (board.Tile, int) {
	startContent := logger.ParseStartEntryContent(entry.Content)

	placedStartTile := elements.ToPlacedTile(startContent.StartingTile)

	tile := board.NewTile(position.New(0, 0), factory.FieldColor, rl.Red)

	for _, f := range placedStartTile.Features {
		if f.FeatureType != engineFeature.Field {
			tile.AddFeature(parseFeatures(f))
		} else {
			tile.AddFeatureBelowOthers(parseFeatures(f))
		}
	}

	return tile, startContent.PlayerCount
}

func ParsePlaceTileEntry(entry logger.Entry) board.Tile {
	placedTileContent := logger.ParsePlaceTileEntryContent(entry.Content)

	tile := board.NewTile(placedTileContent.Move.Position, factory.FieldColor, rl.LightGray)

	for _, f := range placedTileContent.Move.Features {
		if f.FeatureType != engineFeature.Field {
			tile.AddFeature(parseFeatures(f))
		} else {
			tile.AddFeatureBelowOthers(parseFeatures(f))
		}
	}

	return tile
}

func ParseScoreEntry(entry logger.Entry) elements.ScoreReport {
	scoreContent := logger.ParseScoreEntryContent(entry.Content)

	return scoreContent.Scores
}
