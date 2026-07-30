package main

import (
	"os"
	"slices"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	dl := len(d)
	if dl != 52 {
		t.Errorf("Expected deck length of 52, but got %v", dl)
	}
	if d[0] != "A♠" {
		t.Errorf("Expected first card to be A♠, but got %v", d[0])
	}
}

func TestDealHand(t *testing.T) {
	d := newDeck()
	h, r := d.dealHand(5)
	dl := len(d)
	hl, rl := len(h), len(r)

	if dl != 52 {
		t.Errorf("Expected deck length of 52, but got %v", dl)
	}
	if hl != 5 {
		t.Errorf("Expected hand length of 5, but got %v", hl)
	}
	if rl != dl-hl {
		t.Errorf("After dealing a hand, expected the rest of deck to be of size %v, but go %v", dl-hl, rl)
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	filename := "test.deck"
	d1 := newDeck()
	d1.saveToFile(filename)
	d2, _ := newDeckFromFile(filename)
	if !slices.Equal(d1, d2) {
		t.Errorf("Expected that the new and the read from file decks are equal, but they are not:\n%v\n%v", d1, d2)
	}
	os.Remove(filename)
}
