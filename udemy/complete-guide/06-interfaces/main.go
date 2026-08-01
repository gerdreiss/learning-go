package main

import (
	"fmt"
	"time"
)

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

	person := Person{
		"John",
		"Doe",
		time.Date(1990, time.June, 15, 0, 0, 0, 0, time.UTC),
	}

	printFullNameAndAge(person)

	sum := Sum(map[string]float64{
		"first":  35.98,
		"second": 26,
	})
	fmt.Println(sum)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func printFullNameAndAge(info PersonInfo) {
	fmt.Println(info.getFullName(), info.getAge())
}
