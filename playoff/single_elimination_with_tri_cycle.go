// Copyright 2022 Team 254. All Rights Reserved.
// Author: pat@patfairbank.com (Patrick Fairbank)
//
// Defines the tournament structure for a single-elimination, best-of-three bracket.

package playoff

import (
	"fmt"
)

// Creates a single-elimination bracket containing only the required matchups for the given number of alliances, and
// returns the root matchup comprising the tournament finals along with scheduled breaks.
func newSingleEliminationWithTriCycleBracket(numAlliances int) (*Matchup, *TriCycle, []breakSpec, error) {
	if numAlliances != 5 {
		return nil, nil, nil, fmt.Errorf("single-elimination bracket must have at least 2 alliances")
	}

	// Define semifinal matches.
	sf1 := Matchup{
		id:                 "SF1",
		NumWinsToAdvance:   2,
		redAllianceSource:  allianceSelectionSource{1},
		blueAllianceSource: allianceSelectionSource{4},
		matchSpecs: []*matchSpec{
			newSingleEliminationMatch("Semifinal", "SF", 1, 1, 37),
			newSingleEliminationMatch("Semifinal", "SF", 1, 2, 39),
			newSingleEliminationMatch("Semifinal", "SF", 1, 3, 41),
		},
	}
	sf2 := Matchup{
		id:                 "SF2",
		NumWinsToAdvance:   2,
		redAllianceSource:  allianceSelectionSource{2},
		blueAllianceSource: allianceSelectionSource{3},
		matchSpecs: []*matchSpec{
			newSingleEliminationMatch("Semifinal", "SF", 2, 1, 38),
			newSingleEliminationMatch("Semifinal", "SF", 2, 2, 40),
			newSingleEliminationMatch("Semifinal", "SF", 2, 3, 42),
		},
	}

	tricycle := TriCycle{
		id: "TRI",
		alliances: [3]allianceSource{
			newSingleEliminationLooseAllianceSource(&sf1, numAlliances),
			newSingleEliminationLooseAllianceSource(&sf2, numAlliances),
			allianceSelectionSource{5},
		},
		matchSpecs: []*matchSpec{
			newSingleEliminationTriCycleMatch("TriCycle", "TRI", 1, 1, 43),
			newSingleEliminationTriCycleMatch("TriCycle", "TRI", 2, 1, 44),
			newSingleEliminationTriCycleMatch("TriCycle", "TRI", 3, 1, 45),
			newSingleEliminationTriCycleMatch("TriCycle", "TRI", 1, 2, 46),
			newSingleEliminationTriCycleMatch("TriCycle", "TRI", 2, 2, 47),
			newSingleEliminationTriCycleMatch("TriCycle", "TRI", 3, 3, 48),
		},
	}

	// Define final matches.
	final := Matchup{
		id:                 "F",
		NumWinsToAdvance:   2,
		redAllianceSource:  newSingleEliminationAllianceSource(&sf1, numAlliances),
		blueAllianceSource: newSingleEliminationAllianceSource(&sf2, numAlliances),
		matchSpecs:         newFinalMatches(49),
	}

	// Define scheduled breaks.
	var breakSpecs []breakSpec
	breakSpecs = append(breakSpecs, breakSpec{50, 480, "Field Break"})
	breakSpecs = append(breakSpecs, breakSpec{51, 480, "Field Break"})

	return &final, &tricycle, breakSpecs, nil
}

// Helper method to create an allianceSource while pruning any unnecessary matchups due to the number of alliances.
func newSingleEliminationLooseAllianceSource(matchup *Matchup, numAlliances int) allianceSource {
	redAllianceId := matchup.redAllianceSource.AllianceId()
	blueAllianceId := matchup.blueAllianceSource.AllianceId()

	if blueAllianceId > redAllianceId && blueAllianceId > numAlliances {
		return matchup.blueAllianceSource
	}
	if redAllianceId > blueAllianceId && redAllianceId > numAlliances {
		return matchup.redAllianceSource
	}
	return matchupSource{matchup: matchup, useWinner: false}
}

func newSingleEliminationTriCycleMatch(longRoundName, shortRoundName string, setNumber, matchNumber, order int) *matchSpec {
	return &matchSpec{
		longName:            fmt.Sprintf("%s %d-%d", longRoundName, setNumber, matchNumber),
		shortName:           fmt.Sprintf("%s%d-%d", shortRoundName, setNumber, matchNumber),
		order:               order,
		durationSec:         600,
		useTiebreakCriteria: true,
	}
}
