package main

import (
	"fmt"
	"net/http"

	"github.com/gerdreiss/go-course/pkg/handlers"
	"github.com/gerdreiss/go-course/pkg/utils"
)

const addr = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/divide/{x}/{y}", utils.Divide)

	fmt.Printf("Starting server on port %s...\n", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("WTF!", err)
	}
}
