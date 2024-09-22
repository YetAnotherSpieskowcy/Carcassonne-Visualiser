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
	if f.FeatureType == engineFeature.Monastery {
		return factory.Monastery(f)
	} else if f.FeatureType == engineFeature.Road {
		return factory.Road(f.Sides)
	} else {
		return factory.City(f.Sides, f.ModifierType == engineModifier.Shield)
	}
}

func ParseStartEntry(entry logger.Entry) (board.Tile, int) {
	startContent := logger.ParseStartEntryContent(entry.Content)

	placedStartTile := elements.ToPlacedTile(startContent.StartingTile)

	tile := board.NewTile(position.New(0, 0), rl.DarkGreen, rl.Red)

	for _, f := range placedStartTile.Features {
		if f.FeatureType != engineFeature.Field {
			tile.AddFeature(parseFeatures(f))
		}
	}

	return tile, startContent.PlayerCount
}

func ParsePlaceTileEntry(entry logger.Entry) board.Tile {
	placedTileContent := logger.ParsePlaceTileEntryContent(entry.Content)

	tile := board.NewTile(placedTileContent.Move.Position, rl.DarkGreen, rl.LightGray)

	for _, f := range placedTileContent.Move.Features {
		if f.FeatureType != engineFeature.Field {
			tile.AddFeature(parseFeatures(f))
		}
	}

	return tile
}

func ParseScoreEntry(entry logger.Entry) elements.ScoreReport {
	scoreContent := logger.ParseScoreEntryContent(entry.Content)

	return scoreContent.Scores
}
