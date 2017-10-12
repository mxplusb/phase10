package phases

import "github.com/mxplusb/phaseten/cards"

type PhaseNumber int

const (
	One   PhaseNumber = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
)

type Phase struct {
	CurrentPhase PhaseNumber
	PhaseType    interface{}
}

type Run struct {
	Length int
}

type PhaseSet struct {
	Sets     []*Set
	NumSets int
}

// NewPhaseSet returns a new
func NewPhaseSet(i int) *PhaseSet {
	return &PhaseSet{NumSets: i}
}

type Set struct {
	Length int
}

func (s *Set) Validate(num int, c []*cards.Card) bool {
	return false
}
