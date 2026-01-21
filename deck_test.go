package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be 'Ace of Spades', but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card to be 'Four of Clubs', but got %v", d[len(d)-1])
	}
}

func TestDeal(t *testing.T) {
	d := newDeck()
	handSize := 5
	hand, remainingDeck := deal(d, handSize)

	if len(hand) != handSize {
		t.Errorf("Expected hand size of %d, but got %d", handSize, len(hand))
	}

	if len(remainingDeck) != len(d)-handSize {
		t.Errorf("Expected remaining deck size of %d, but got %d", len(d)-handSize, len(remainingDeck))
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	filename := "_decktesting"

	// Clean up test file after test
	defer os.Remove(filename)

	d := newDeck()
	err := d.saveToFileOs(filename)
	if err != nil {
		t.Errorf("Error saving deck to file: %v", err)
	}

	loadedDeck := newDeckFromFile(filename)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected loaded 16 cards in deck, but got %v", len(loadedDeck))
	}

	if len(loadedDeck) != len(d) {
		t.Errorf("Expected loaded deck size of %d, but got %d", len(d), len(loadedDeck))
	}

	for i, card := range d {
		if loadedDeck[i] != card {
			t.Errorf("Expected card %s at position %d, but got %s", card, i, loadedDeck[i])
		}
	}
}

func TestShuffle(t *testing.T) {
	d := newDeck()
	originalDeck := make(deck, len(d))
	copy(originalDeck, d)

	d.shuffle()

	sameOrder := true
	for i := range d {
		if d[i] != originalDeck[i] {
			sameOrder = false
			break
		}
	}

	if sameOrder {
		t.Errorf("Expected deck to be shuffled, but order is the same")
	}
}
