package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://stackoverflow.com",
		"http://duckduckgo.com",
		"http://golang.org",
		"http://github.com",
		"http://gitlab.com",
	}

	// create a channel
	channel := make(chan string)

	for _, link := range links {
		// pass the channel to the function
		go checkLink(link, channel)
	}

	for link := range channel {
		go func(link string, pause time.Duration) {
			// wait for the duration of pause
			time.Sleep(pause)
			// receive messages from the channel and check again
			checkLink(link, channel)
		}(link, 3*time.Second)
	}
}

func checkLink(link string, channel chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " is not reachable")
	} else {
		fmt.Println(link, " is up")
	}
	channel <- link
}
