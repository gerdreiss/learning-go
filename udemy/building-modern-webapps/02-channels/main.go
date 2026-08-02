package main

import (
	"log"

	"github.com/gerdreiss/channels/helpers"
)

func CalculateValue(ic chan int) {
	rn := helpers.RandomNumber(1000)
	ic <- rn
}

func main() {
	ic := make(chan int)
	defer close(ic)

	go CalculateValue(ic)

	num := <-ic

	log.Println(num)
}
