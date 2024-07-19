package playoff

type TriCycle struct {
	id         string
	alliances  [3]allianceSource
	matchSpecs []*matchSpec
}

func (triCycle *TriCycle) Id() string {
	return triCycle.id
}

func (triCycle *TriCycle) MatchSpecs() []*matchSpec {
	return triCycle.matchSpecs
}

func (triCycle *TriCycle) update(playoffMatchResults map[int]playoffMatchResult) {

}

func (triCycle *TriCycle) traverse(visitFunction func(MatchGroup) error) error {
	if err := visitFunction(triCycle); err != nil {
		return err
	}
	for _, alliance := range triCycle.alliances {
		if err := alliance.traverse(visitFunction); err != nil {
			return err
		}
	}
	return nil
}
