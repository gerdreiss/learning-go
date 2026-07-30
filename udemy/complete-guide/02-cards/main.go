package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"slices"
)

func clearScreen() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default: // Linux, macOS, Unix
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()

	fmt.Println()
}

func main() {
	clearScreen()

	filename := "deck.txt"

	deck := newDeck()
	deck.saveToFile(filename)
	read := newDeckFromFile(filename)

	if !slices.Equal(deck, read) {
		fmt.Println("something went wrong when reading the deck from ", filename, ": deck =", deck, "read =", read)
		os.Exit(1)
	}

	hand, _ := dealHand(deck, 5)
	fmt.Println("\nDealt hand:")
	hand.print()
}
