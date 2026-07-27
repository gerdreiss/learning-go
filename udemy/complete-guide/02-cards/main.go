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

	fmt.Println("==================================================")

	deck := generateDeck()
	fmt.Println("The full deck of cards:")
	deck.print()

	fmt.Println("==================================================")

	hand := selectRandomHand(deck, 3)
	fmt.Println("Randomly selected hand:")
	for i, card := range hand {
		fmt.Println(i, card)
	}
}
