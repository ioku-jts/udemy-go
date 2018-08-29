package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()
	if len(deck) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(deck))
	}

	if deck[0] != "A of Spades" {
		t.Errorf("Expected first card to be 'Ace of Spades', but got %v", deck[0])
	}

	if deck[len(deck)-1] != "K of Diamonds" {
		t.Errorf("Expected first card to be 'King of Diamonds', but got %v", deck[len(deck)-1])
	}
}

func TestSaveAndReadFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")
	newD := readFromFile("_decktesting")

	if len(newD) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(deck))
	}

	os.Remove("_decktesting")
}
