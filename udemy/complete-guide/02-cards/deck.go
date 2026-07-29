package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type deck []string
type hand []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (h hand) print() {
	for i, card := range h {
		fmt.Println(i, card)
	}
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(strings.Join(d, "\n")), 0666)
}

func newDeck() deck {
	suits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	ranks := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	deck := make(deck, 0, 52)
	for _, suit := range suits {
		for _, rank := range ranks {
			deck = append(deck, rank+" of "+suit)
		}
	}
	return deck
}

func dealHand(d deck, size int) (hand, deck) {
	shuffled := shuffleDeck(d)

	// Return the first 'size' cards and the rest
	return shuffled[:size], shuffled[size:]
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
