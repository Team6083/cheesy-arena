// Copyright 2023 Team 254. All Rights Reserved.
// Author: pat@patfairbank.com (Patrick Fairbank)
//
// Model representing the instantaneous score of a match.

package game

type Score struct {
	Cubes           [5]int
	CubeBonus       int
	EndgameStatuses [3]EndgameStatus
	Fouls           []Foul
	PlayoffDq       bool
}

// Game-specific constants that cannot be changed by the user.
const ()

// Game-specific settings that can be changed by the user.

// Represents the state of a robot at the end of the match.
type EndgameStatus int

const (
	EndgameNone EndgameStatus = iota
	EndgameParked
)

// Calculates and returns the summary fields used for ranking and display.
func (score *Score) Summarize(opponentScore *Score) *ScoreSummary {
	summary := new(ScoreSummary)

	// Leave the score at zero if the alliance was disqualified.
	if score.PlayoffDq {
		return summary
	}

	// Calculate Cube points.
	summary.CubePoints = 0
	for i, cube := range score.Cubes {
		oneCube := 10 + 5*i
		summary.CubePoints += oneCube * cube
	}
	summary.CubePoints += 5 * score.CubeBonus

	// Calculate endgame points.
	for _, status := range score.EndgameStatuses {
		switch status {
		case EndgameParked:
			summary.ParkPoints += 1
		default:
		}
	}

	summary.MatchPoints = summary.CubePoints + summary.ParkPoints

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
	if score.CubeBonus != other.CubeBonus ||
		score.EndgameStatuses != other.EndgameStatuses ||
		score.PlayoffDq != other.PlayoffDq ||
		len(score.Fouls) != len(other.Fouls) ||
		len(score.Cubes) != len(other.Cubes) {
		return false
	}

	for i, foul := range score.Fouls {
		if foul != other.Fouls[i] {
			return false
		}
	}

	for i, cube := range score.Cubes {
		if cube != other.Cubes[i] {
			return false
		}
	}

	return true
}
