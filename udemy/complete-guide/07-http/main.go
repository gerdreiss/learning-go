package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://duckduckgo.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Println("Request to DDG	returned:")
	fmt.Printf("Response status: %v\n", resp.Status)

	// writeToFileBadly(resp)

	writeToFile(resp)
}

func writeToFile(resp *http.Response) {
	dst, err := os.Create("duckduckgo.html")
	if err != nil {
		fmt.Printf("Unexpected error occured: %v\n", err)
		os.Exit(1)
	}
	defer dst.Close()

	io.Copy(dst, resp.Body)
}

func writeToFileBadly(resp *http.Response) {
	bs := make([]byte, 9999)
	n, err := resp.Body.Read(bs)
	if err != nil {
		fmt.Printf("Unexpected error occured: %v\n", err)
		os.Exit(1)
	}
	if n == 0 {
		fmt.Println("Received empty response")
		os.Exit(0)
	}
	if err := os.WriteFile("ddg.html", bs, 0666); err != nil {
		fmt.Printf("Writing DDG response to file failed: %v\n", err)
		os.Exit(1)
	}
}
