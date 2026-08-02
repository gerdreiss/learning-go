package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello, world!")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Bytes written: %d\n\n", n)
	},
	)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("WTF!", err)
	}
}
