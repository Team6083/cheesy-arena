// Copyright 2023 Team 254. All Rights Reserved.
// Author: pat@patfairbank.com (Patrick Fairbank)
//
// Model representing the instantaneous score of a match.

package game

type Score struct {
	CollectionCubes [5]int
	CollectionBonus int
	PushCubes       [2]int
	PushBonus       int
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

	// Calculate Push points.
	summary.PushPoints = 0
	for i, cubes := range score.PushCubes {
		pointsPerCube := 1 + 1*i
		summary.PushPoints += pointsPerCube * cubes
	}
	summary.PushPoints += 4 * score.PushBonus

	// Calculate Collection points.
	summary.CollectionPoints = 0
	for i, cubes := range score.CollectionCubes {
		pointsPerCube := 10 + 5*i
		summary.CollectionPoints += pointsPerCube * cubes
	}
	summary.CollectionPoints += 5 * score.CollectionBonus

	// Calculate endgame points.
	for _, status := range score.EndgameStatuses {
		switch status {
		case EndgameParked:
			summary.ParkPoints += 5
		default:
		}
	}

	summary.MatchPoints = summary.PushPoints + summary.CollectionPoints + summary.ParkPoints

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
	if score.CollectionBonus != other.CollectionBonus ||
		score.EndgameStatuses != other.EndgameStatuses ||
		score.PlayoffDq != other.PlayoffDq ||
		len(score.Fouls) != len(other.Fouls) {
		return false
	}

	for i, foul := range score.Fouls {
		if foul != other.Fouls[i] {
			return false
		}
	}

	for i, col := range score.CollectionCubes {
		if col != other.CollectionCubes[i] {
			return false
		}
	}

	for i, push := range score.PushCubes {
		if push != other.PushCubes[i] {
			return false
		}
	}

	return true
}
