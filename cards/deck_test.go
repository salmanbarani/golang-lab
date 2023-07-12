package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expedted deck length to be 16, but receive %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expedted deck length to be 'Ace of Spades', but receive %v", len(d))
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expedted deck length to be 'Four of Clubs', but receive %v", len(d))
	}
}

func TestSaveToFileAndnewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expedted deck length to be 16, but receive %v", len(loadedDeck))
	}

	os.Remove("_decktesting") // Always must be removed after testing.

}
