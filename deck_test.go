package main

import (
	"os"
	"strconv"
	"testing"
)

func AceStrConv(d Deck) string {
	cardStringSlice := ""
	cardStringSlice += d[0].rank.cardname + d[0].suit + " " + strconv.Itoa(d[0].rank.value)
	return cardStringSlice
}

func KingStrConv(d Deck) string {
	cardStringSlice := ""
	cardStringSlice += d[51].rank.cardname + d[51].suit + " " + strconv.Itoa(d[51].rank.value)
	return cardStringSlice
}

func TestNewDeck(t *testing.T) {
	d := newDeck()
	ace := AceStrConv(d)
	king := KingStrConv(d)

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if ace != "Ace of Spades 1" {
		t.Errorf("Expected 'Ace of Spaces 1' but failed to convert")
	}

	if king != "King of Hearts 13" {
		t.Errorf("Expected 'Ace of Spaces 1' but failed to convert")
	}
}

func TestSaveToDeckandNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := bytesliceToDeck(deckFromFile("_decktesting"))

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
