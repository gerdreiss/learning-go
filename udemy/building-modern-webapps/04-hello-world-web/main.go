package main

import (
	"fmt"
	"net/http"
)

const addr = ":8080"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Printf("Starting server on port %s...\n", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("WTF!", err)
	}
}
