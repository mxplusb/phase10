package phases

import "github.com/mxplusb/phaseten/cards"

type PhaseOne struct {
	ph      Phase
	NumSets int
}

func NewPhaseOne() *PhaseOne {
	return &PhaseOne{
		NumSets: 2,
		ph: Phase{
			PhaseType: Set{},
		},
	}
}

//ValidateSets validates two sets of three. It will fail if two sets of the same number are used.
func (p *PhaseOne) Validate(x, y []*cards.Card) bool {
	// first Set
	x1, x2, x3 := x[0], x[1], x[2]
	// second set.
	y1, y2, y3 := y[0], y[1], y[2]

	if x1 != x2 && x2 != x3 {
		return false
	}
	if y1 != y2 && y2 != y3 {
		return false
	}
	if y1 == x1 {
		return false
	}
	return true
}
