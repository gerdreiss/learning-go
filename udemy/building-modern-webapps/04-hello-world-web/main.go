package main

import (
	"errors"
	"fmt"
	"net/http"
)

const addr = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "Hello, world!")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Bytes written: %d\n\n", n)
}

func About(w http.ResponseWriter, r *http.Request) {
	res, err := divide(2, 2)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
	}
	fmt.Fprintf(w, "This is the about page and 2/2 is %d\n", res)
}

func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("division by zero")
	}
	return x / y, nil
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Starting server on port %s...\n", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("WTF!", err)
	}
}
