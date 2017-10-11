package phases

type PhaseNumber int

const (
	One PhaseNumber = iota
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

type Set struct {
	Length int
}
