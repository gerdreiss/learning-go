package main

import (
	"fmt"
	"math/rand"
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

func generateDeck() deck {
	suits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	ranks := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}

	deck := make(deck, 0, 52)
	for _, suit := range suits {
		for _, rank := range ranks {
			deck = append(deck, rank+" of "+suit)
		}
	}
	return deck
}

func selectRandomHand(d deck, count int) hand {
	// Create a copy of the deck to avoid modifying the original
	shuffled := make([]string, len(d))
	copy(shuffled, d)

	// Shuffle the deck using Fisher-Yates algorithm
	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})

	// Return the first 'count' cards
	return shuffled[:count]
}
