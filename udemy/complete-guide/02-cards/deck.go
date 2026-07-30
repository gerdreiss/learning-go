package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type deck []string
type hand []string

func newDeck() deck {
	suits := []string{"♠", "♥", "♦", "♣"}
	ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	deck := make(deck, 0, len(suits)*len(ranks))
	for _, suit := range suits {
		for _, rank := range ranks {
			deck = append(deck, rank+suit)
		}
	}
	return deck
}

// newDeckFromFile reads a deck from a file
func newDeckFromFile(filename string) (deck, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
	}

	// Handle empty or whitespace-only files
	content := strings.TrimSpace(string(bytes))
	if content == "" {
		return nil, fmt.Errorf("file %s is empty", filename)
	}

	return strings.Split(content, "\n"), nil
}

func shuffleDeck(d deck) []string {
	// Create a copy of the deck to avoid modifying the original
	shuffled := make([]string, len(d))
	copy(shuffled, d)

	// Shuffle the deck using Fisher-Yates algorithm
	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})

	return shuffled
}

// print displays all cards in the deck
func (d deck) print() {
	for i, card := range d {
		fmt.Printf("%3d: %s\n", i, card)
	}
}

func (h hand) print() {
	for i, card := range h {
		fmt.Println(i, card)
	}
}

// saveToFile writes the deck to a file
func (d deck) saveToFile(filename string) error {
	if len(d) == 0 {
		return fmt.Errorf("cannot save empty deck")
	}
	return os.WriteFile(filename, []byte(strings.Join(d, "\n")), 0644)
}

func (d deck) dealHand(size int) (hand, deck) {
	// shuffled := shuffleDeck(d)
	// return shuffled[:size], shuffled[size:]

	// shuffle deck
	d.shuffle()

	// Return the first 'size' cards and the rest
	return hand(d[:size]), d[size:]
}

func (d deck) shuffle() {
	for i := range d {
		newPos := rand.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i]
	}
}
