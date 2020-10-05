package main

import (
	"os"
	"strings"
	"testing"
)

func TestNewDeck(t *testing.T) {
	cards := newDeck()
	if len(cards) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(cards))
	}

	if strings.TrimSpace(cards[0]) != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", cards[0])
	}

	lastCard := strings.TrimSpace(cards[len(cards)-1])
	if lastCard != "King of Clubs" {
		t.Errorf("Expected last card of King of Clubs, but got %v", lastCard)
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	fileName := "_decktesting.txt"
	os.Remove(fileName)

	deck := newDeck()
	deck.saveToFile(fileName)

	loadedDeck := newDeckFromFile(fileName)

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(loadedDeck))
	}
	os.Remove(fileName)
}
