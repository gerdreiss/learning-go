package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage <command> filename")
		os.Exit(1)
	}

	filename := os.Args[1]
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		fmt.Errorf("failed to read file %s: %w", filename, err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}
