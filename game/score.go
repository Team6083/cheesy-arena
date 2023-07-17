// Copyright 2023 Team 254. All Rights Reserved.
// Author: pat@patfairbank.com (Patrick Fairbank)
//
// Model representing the instantaneous score of a match.

package game

type Score struct {
	PowerCell    int
	Block        int
	CubeStatus   bool
	ParkStatuses [3]bool
	Fouls        []Foul
	PlayoffDq    bool
}

var SustainabilityBonusLinkThresholdWithoutCoop = 7
var SustainabilityBonusLinkThresholdWithCoop = 6
var ActivationBonusPointThreshold = 26

// Calculates and returns the summary fields used for ranking and display.
func (score *Score) Summarize(opponentScore *Score) *ScoreSummary {
	summary := new(ScoreSummary)

	// Leave the score at zero if the alliance was disqualified.
	if score.PlayoffDq {
		return summary
	}

	// Calculate teleoperated period points.
	summary.PowerCellPoints = score.PowerCell * 10
	summary.BlockPoints = score.Block * 5

	parkPoints := 0
	for i := 0; i < 3; i++ {
		if score.ParkStatuses[i] {
			parkPoints += 10
		}
	}

	cubePoints := 0
	if score.CubeStatus {
		cubePoints += 30
	}

	summary.ParkingPoints = parkPoints
	summary.CubePoints = cubePoints
	summary.EndgamePoints = parkPoints + cubePoints
	summary.MatchPoints = summary.PowerCellPoints + summary.BlockPoints + summary.EndgamePoints

	// Calculate penalty points.
	for _, foul := range opponentScore.Fouls {
		summary.FoulPoints += foul.PointValue()
		// Store the number of tech fouls since it is used to break ties in playoffs.
		if foul.IsTechnical {
			summary.NumOpponentTechFouls++
		}
	}

	summary.Score = summary.MatchPoints + summary.FoulPoints

	return summary
}

// Returns true if and only if all fields of the two scores are equal.
func (score *Score) Equals(other *Score) bool {
	if score.ParkStatuses != other.ParkStatuses ||
		score.CubeStatus != other.CubeStatus ||
		score.PowerCell != other.PowerCell ||
		score.Block != other.Block ||
		score.PlayoffDq != other.PlayoffDq ||
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
