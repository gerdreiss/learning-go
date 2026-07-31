package main

import (
	"fmt"
	"strings"
)

func printMap(colors map[string]string) {
	// Print header
	fmt.Printf("%-12s %-10s\n", "Color", "Hex Code")
	fmt.Println(strings.Repeat("-", 27))

	// Print each entry with fixed-width columns
	for key, value := range colors {
		fmt.Printf("%-12s %-10s\n", key, value)
	}
}

func main() {

	//colors := make(map[string]string)
	colors := map[string]string{
		"black":   "#000000",
		"white":   "#FFFFFF",
		"red":     "#FF0000",
		"green":   "#00FF00",
		"blue":    "#0000FF",
		"yellow":  "#FFFF00",
		"cyan":    "#00FFFF",
		"magenta": "#FF00FF",
		"silver":  "#C0C0C0",
		"gray":    "#808080",
		"maroon":  "#800000",
		"olive":   "#808000",
		"purple":  "#800080",
		"teal":    "#008080",
		"navy":    "#000080",
		"orange":  "#FFA500",
	}

	colors["tomato"] = "#FF6347"
	delete(colors, "tomato")
	fmt.Println(colors["tomato"])

	printMap(colors)
}
