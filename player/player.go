package player

import (
	"github.com/mxplusb/phaseten/cards"
	"github.com/mxplusb/phaseten/phases"
)

// Player represents a player.
type Player struct {
	Name string
	Hand         []*cards.Card
	CurrentPhase *phases.Phase
}

// NewPlayer makes a new Player with name n.
func NewPlayer(n string) *Player {
	return &Player{Name: n}
}

// InitialDraw draws 10 cards from the pile and places the Cards in the Player's Hand.
func (p *Player) InitialDraw(d *cards.Deck) {
	for i := 0; i < 10; i++ {
		p.Hand = append(p.Hand, d.Draw())
	}
}

// Draw retrieves one card from the deck and places it in the Player's hand.
func (p *Player) Draw(d *cards.Deck) {
	p.Hand = append(p.Hand, d.Draw())
}