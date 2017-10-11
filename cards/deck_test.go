package cards

import (
	"testing"
	"reflect"
)

func TestBuildDeck(t *testing.T) {
	deck := BuildDeck()
	if len(deck.Cards) != 108 {
		t.Logf("deck does not have 108 cards, it has %d", len(deck.Cards))
		t.Fail()
	}
}

func TestValiadateDeckColours(t *testing.T) {
	deck := BuildDeck()

	var (
		reds int
		blues int
		greens int
		golds int
	)

	for card := range deck.Cards {
		if deck.Cards[card].Color == Red {
			reds++
		}
		if deck.Cards[card].Color == Blue {
			blues++
		}
		if deck.Cards[card].Color == Gold {
			golds++
		}
		if deck.Cards[card].Color == Green {
			greens++
		}
	}
	if (reds != 24) || (blues != 24) || (golds != 24) || (greens != 24) {
		t.Log("there is an incorrect number of coloured cards.")
		t.Logf("there are %d reds, %d blues, %d golds, and %d greens", reds, blues, golds, greens)
		t.Fail()
	}
}

func TestValidateDeckAlternates(t *testing.T) {
	deck := BuildDeck()

	var (
		wilds int
		skips int
	)

	for card := range deck.Cards {
		if deck.Cards[card].Wild {
			wilds++
		}
		if deck.Cards[card].Skip {
			skips++
		}
	}

	if wilds != 8 {
		t.Logf("expecting %d wild cards, got %d", 8, wilds)
		t.Fail()
	}

	if skips != 4 {
		t.Logf("expecting %d skip cards, got %d", 4, skips)
	}
}

func TestDeckShuffle(t *testing.T) {
	deck := BuildDeck()
	var existingOrder []*Card
	copy(existingOrder, deck.Cards)
	deck.Shuffle()
	if reflect.DeepEqual(existingOrder, deck.Cards) {
		t.Logf("cards did not shuffle!")
		t.Fail()
	}
}