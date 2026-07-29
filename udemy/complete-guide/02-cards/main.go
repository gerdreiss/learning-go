package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
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

	deck := newDeck()
	deck.saveToFile("deck.txt")

	hand, rest := dealHand(deck, 5)
	fmt.Println("\nDealt hand:")
	hand.print()
	fmt.Println("\nRest of the deck:")
	rest.print()
}
