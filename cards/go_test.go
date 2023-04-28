package main

import (
	"os"
	"testing"
)

func TestNewDek(t *testing.T) {
	d := newDeck()

	lenght := len(d)

	if lenght != 16 {
		t.Errorf("Expected deck lenght of 16 but got %d", lenght)
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewTestFromFile(t *testing.T) {
	fileName := "_decktesting"

	os.Remove(fileName)

	deck := newDeck()

	deck.saveToFile(fileName)

	loadedDeck := newDeckFromFile(fileName)

	lenght := len(loadedDeck)

	if lenght != 16 {
		t.Errorf("Expected deck lenght of 16 but got %d", lenght)
	}

	os.Remove(fileName)
}
