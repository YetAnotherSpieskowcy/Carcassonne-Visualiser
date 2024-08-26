package addons

import (
	"fmt"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type ScoreInfo struct {
	scoresToPrint    map[elements.ID]uint32
	scoreReports     map[uint32]elements.ScoreReport
	nextScoreReports map[uint32]elements.ScoreReport

	//playerScores     map[uint32]map[elements.ID]uint32
	//nextPlayerScores map[uint32]map[elements.ID]uint32
	position rl.Vector2
}

func NewScoreInfo(playerCount int, pos rl.Vector2) ScoreInfo {
	scoresToPrint := make(map[elements.ID]uint32)
	i := 1
	for i <= playerCount {
		scoresToPrint[elements.ID(i)] = 0
		i++
	}

	scoreReport := elements.NewScoreReport()
	scoreReport.ReceivedPoints = scoresToPrint

	scoreReports := make(map[uint32]elements.ScoreReport)
	scoreReports[0] = scoreReport

	return ScoreInfo{
		scoresToPrint:    scoresToPrint,
		scoreReports:     scoreReports,
		nextScoreReports: make(map[uint32]elements.ScoreReport),
		position:         pos,
	}
}

func (scoreInfo *ScoreInfo) PreviousScores(move uint32) {
	scoreReport, ok := scoreInfo.scoreReports[move]
	if ok {
		for id, score := range scoreReport.ReceivedPoints {
			scoreInfo.scoresToPrint[id] -= score
		}
		scoreInfo.nextScoreReports[move] = scoreReport
		delete(scoreInfo.scoreReports, move)
	}
}

func (scoreInfo *ScoreInfo) NextScores(move uint32) {
	scoreReport, ok := scoreInfo.nextScoreReports[move]
	if ok {
		for id, score := range scoreReport.ReceivedPoints {
			scoreInfo.scoresToPrint[id] += score
		}
		scoreInfo.scoreReports[move] = scoreReport
		delete(scoreInfo.nextScoreReports, move)
	}
}

func (scoreInfo *ScoreInfo) UpdateScores(scoreReport elements.ScoreReport, move uint32) {
	scoreInfo.scoreReports[move] = scoreReport
	for id, score := range scoreReport.ReceivedPoints {
		scoreInfo.scoresToPrint[id] += score
	}
}

func (scoreInfo ScoreInfo) Show(move uint32) {
	position := int32(scoreInfo.position.Y)
	// show scores
	for id, score := range scoreInfo.scoresToPrint {
		text := "Player " + fmt.Sprint(id) + ": " + fmt.Sprint(score)
		rl.DrawText(text, int32(scoreInfo.position.X), position+int32(id*20), 20, rl.Black)
	}

	// show returned meeples
	scoreReport, ok := scoreInfo.scoreReports[move]

	if ok && move > 0 {
		rl.DrawText("Returned meeples:", int32(scoreInfo.position.X), position+80, 20, rl.Black)
		position += 100

		// display returned meeples in ascending player ID order. Without this the order is random in every frame, and the text blinks a lot
		maxPlayerID := elements.ID(0)
		for id := range scoreReport.ReturnedMeeples {
			if id > maxPlayerID {
				maxPlayerID = id
			}
		}
		for id := range maxPlayerID + 1 {
			returnedMeeples, ok := scoreReport.ReturnedMeeples[id]
			if !ok {
				continue
			}
			text := "Player " + fmt.Sprint(id) + ":"
			rl.DrawText(text, int32(scoreInfo.position.X), position, 20, rl.Black)
			for idx, meeple := range returnedMeeples {
				text := " - " + fmt.Sprint(meeple.Type) + " at x = " + fmt.Sprint(meeple.Position.X()) + ", y = " + fmt.Sprint(meeple.Position.Y())
				rl.DrawText(text, int32(scoreInfo.position.X), position+int32((idx+1)*20), 20, rl.Black)
			}
			position += int32((len(returnedMeeples) + 1) * 20)
		}
	}
}
