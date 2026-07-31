package main

import (
	"fmt"
	"time"
)

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func (englishBot) getGreeting() string {
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

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
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func printFullNameAndAge(info PersonInfo) {
	fmt.Println(info.getFullName(), info.getAge())
}
