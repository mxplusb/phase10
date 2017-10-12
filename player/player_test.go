package player

import (
	"testing"
	"github.com/mxplusb/phaseten/cards"
)

func TestPlayerInitialDraw(t *testing.T) {
	deck := cards.BuildDeck()
	deck.Shuffle()

	testPlayer := NewPlayer("test")
	testPlayer.InitialDraw(deck)

	if len(testPlayer.Hand) != 10 {
		t.Logf("expected player to have 10 cards, got %d", len(testPlayer.Hand))
		t.Fail()
	}

	if len(deck.Cards) != 98 {
		t.Logf("expected 98 cards in deck, got %d", len(deck.Cards))
		t.Fail()
	}
}

func TestPlayerDraw(t *testing.T) {
	deck := cards.BuildDeck()
	deck.Shuffle()

	testPlayer := NewPlayer("test")
	testPlayer.Draw(deck)

	if len(testPlayer.Hand) != 1 {
		t.Logf("expected player to have 10 cards, got %d", len(testPlayer.Hand))
		t.Fail()
	}

	if len(deck.Cards) != 107 {
		t.Logf("expected 98 cards in deck, got %d", len(deck.Cards))
		t.Fail()
	}
}