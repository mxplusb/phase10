package cards

import (
	"math/rand"
)

type Color int

// Colors. Wilds and Skips are 10.
const (
	Red   Color = iota
	Green
	Gold
	Blue
)

// Wilds
type Wilder bool

// Skips
type Skipper bool

// One card
type Card struct {
	Value int
	Color Color
	Wild  Wilder
	Skip  Skipper
}

// Phase 10 deck
type Deck struct {
	// All cards.
	Cards []*Card
	// Draw pile for new cards.
	DrawPile []*Card
	// Discard pile for played cards.
	DiscardPile []*Card
}

// BuildDeck will build a brand new, unshuffled deck.
func BuildDeck() *Deck {
	deck := &Deck{}

	// build the numbered cards.
	numberedBuilder := func(d *Deck) {
		for i := 0; i < 12; i++ {
			redCard := &Card{
				Value: i,
				Color: Red,
			}
			blueCard := &Card{
				Value: i,
				Color: Blue,
			}
			greenCard := &Card{
				Value: i,
				Color: Green,
			}
			goldCard := &Card{
				Value: i,
				Color: Gold,
			}
			d.Cards = append(deck.Cards, redCard)
			d.Cards = append(deck.Cards, blueCard)
			d.Cards = append(deck.Cards, greenCard)
			d.Cards = append(deck.Cards, goldCard)
		}
	}

	// generate the numbered deck.
	numberedBuilder(deck)
	numberedBuilder(deck)

	// build our skipped cards.
	for i := 0; i < 4; i++ {
		deck.Cards = append(deck.Cards, &Card{Skip: true, Color: 10})
	}

	// build our wilds.
	for i := 0; i < 8; i++ {
		deck.Cards = append(deck.Cards, &Card{Wild: true, Color: 10})
	}

	return deck
}

// Shuffle implements a Fisher-Yates shuffle.
func (d *Deck) Shuffle() {
	for i := range d.Cards {
		r := rand.Intn(i + 1)
		d.Cards[i], d.Cards[r] = d.Cards[r], d.Cards[i]
	}
}

// Draw removes a Card from the deck and returns it to a Player.
func (d *Deck) Draw() *Card {
	drawn := d.Cards[0]
	d.Cards = append(d.Cards[:0], d.Cards[1:]...)
	return drawn
}
