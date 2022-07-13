// Copyright 2022 Team 254. All Rights Reserved.
// Author: pat@patfairbank.com (Patrick Fairbank)
//
// Model representing the instantaneous score of a match.

package game

type Score struct {
	Pearls int
	Golds  int
	Cube   bool
	Fouls  []Foul
	ElimDq bool
}

type ScoreSummary struct {
	GoldsCount   int
	GoldsPoints  int
	PearlsCount  int
	PearlsPoints int
	CubePoints   int
	CubeAchieved bool
	MatchPoints  int
	FoulPoints   int
	Score        int
}

// var CargoBonusRankingPointThresholdWithoutQuintet = 20
// var CargoBonusRankingPointThresholdWithQuintet = 18
// var HangarBonusRankingPointThreshold = 16

// Represents the state of a robot at the end of the match.
type EndgameStatus int

const (
	EndgameNone EndgameStatus = iota
	EndgameLow
	EndgameMid
	EndgameHigh
	EndgameTraversal
)

// Calculates and returns the summary fields used for ranking and display.
func (score *Score) Summarize(opponentFouls []Foul) *ScoreSummary {
	summary := new(ScoreSummary)

	// Leave the score at zero if the team was disqualified.
	if score.ElimDq {
		return summary
	}

	// Calculate teleoperated period cargo points.
	summary.GoldsCount = score.Golds
	summary.GoldsPoints = 5 * score.Golds
	summary.PearlsCount = score.Pearls
	summary.PearlsPoints = 8 * score.Pearls

	// Calculate endgame points.
	if score.Cube {
		summary.CubePoints = 25
		summary.CubeAchieved = true
	}

	// Calculate bonus ranking points.
	// var cargoBonusRankingPointThreshold int
	// if summary.AutoCargoCount >= QuintetThreshold {
	// 	cargoBonusRankingPointThreshold = CargoBonusRankingPointThresholdWithQuintet
	// 	summary.AutoCargoRemaining = 0
	// 	summary.QuintetAchieved = true
	// } else {
	// 	cargoBonusRankingPointThreshold = CargoBonusRankingPointThresholdWithoutQuintet
	// 	summary.AutoCargoRemaining = QuintetThreshold - summary.AutoCargoCount
	// }
	// if summary.CargoCount >= cargoBonusRankingPointThreshold {
	// 	summary.TeleopCargoRemaining = 0
	// 	summary.CargoBonusRankingPoint = true
	// } else {
	// 	summary.TeleopCargoRemaining = cargoBonusRankingPointThreshold - summary.CargoCount
	// }
	// summary.HangarBonusRankingPoint = summary.HangarPoints >= HangarBonusRankingPointThreshold

	// Calculate penalty points.
	for _, foul := range opponentFouls {
		summary.FoulPoints += foul.PointValue()
	}

	// Check for the opponent fouls that automatically trigger a ranking point.
	// Note: There are no such fouls in the 2022 game; leaving this comment for future years.

	summary.MatchPoints = summary.GoldsPoints + summary.PearlsPoints + summary.CubePoints
	summary.Score = summary.MatchPoints + summary.FoulPoints

	return summary
}

// Returns true if and only if all fields of the two scores are equal.
func (score *Score) Equals(other *Score) bool {
	if score.Cube != other.Cube ||
		score.Golds != other.Golds ||
		score.Pearls != other.Pearls ||
		score.ElimDq != other.ElimDq ||
		len(score.Fouls) != len(other.Fouls) {
		return false
	}

	for i, foul := range score.Fouls {
		if foul != other.Fouls[i] {
			return false
		}
	}

	return true
}
