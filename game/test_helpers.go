// Copyright 2017 Team 254. All Rights Reserved.
// Author: pat@patfairbank.com (Patrick Fairbank)
//
// Helper methods for use in tests in this package and others.

package game

func TestScore1() *Score {
	fouls := []Foul{
		{true, 25, 13},
		{false, 1868, 14},
		{false, 1868, 14},
		{true, 25, 15},
		{true, 25, 15},
		{true, 25, 15},
		{true, 25, 15},
	}
	return &Score{
		EndgameStatuses: [3]EndgameStatus{EndgameParked, EndgameNone},
		Fouls:           fouls,
		PlayoffDq:       false,
	}
}

func TestScore2() *Score {
	return &Score{
		EndgameStatuses: [3]EndgameStatus{},
		Fouls:           []Foul{},
		PlayoffDq:       false,
	}
}

func TestRanking1() *Ranking {
	return &Ranking{254, 1, 0, RankingFields{20, 90, 12, 0.254, 3, 2, 1, 0, 10}}
}

func TestRanking2() *Ranking {
	return &Ranking{1114, 2, 1, RankingFields{18, 625, 23, 0.1114, 1, 3, 2, 0, 10}}
}
