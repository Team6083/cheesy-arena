package playoff

type TriCycle struct {
	id                 string
	redAllianceSource  allianceSource
	blueAllianceSource allianceSource
	matchSpecs         []*matchSpec
	RedAllianceId      int
	BlueAllianceId     int
}

func (triCycle *TriCycle) Id() string {
	return triCycle.id
}

func (triCycle *TriCycle) MatchSpecs() []*matchSpec {
	return triCycle.matchSpecs
}

func (triCycle *TriCycle) update(playoffMatchResults map[int]playoffMatchResult) {
	triCycle.redAllianceSource.update(playoffMatchResults)
	triCycle.blueAllianceSource.update(playoffMatchResults)

	triCycle.RedAllianceId = triCycle.redAllianceSource.AllianceId()
	triCycle.BlueAllianceId = triCycle.blueAllianceSource.AllianceId()

	triCycle.redAllianceSource.setDestination(triCycle)
	triCycle.blueAllianceSource.setDestination(triCycle)

	for _, match := range triCycle.matchSpecs {
		match.redAllianceId = triCycle.RedAllianceId
		match.blueAllianceId = triCycle.BlueAllianceId
	}
}

func (triCycle *TriCycle) traverse(visitFunction func(MatchGroup) error) error {
	if err := visitFunction(triCycle); err != nil {
		return err
	}
	if err := triCycle.redAllianceSource.traverse(visitFunction); err != nil {
		return err
	}
	if err := triCycle.blueAllianceSource.traverse(visitFunction); err != nil {
		return err
	}
	return nil
}
