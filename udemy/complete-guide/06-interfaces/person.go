package main

import (
	"math"
	"time"
)

type Person struct {
	firstName string
	lastName  string
	birthdate time.Time
}

type Age interface {
	getAge() int
}

type FullName interface {
	getFullName() string
}

func (p Person) getAge() int {
	return int(math.Floor(time.Now().In(time.UTC).Sub(p.birthdate).Hours() / (24 * 365)))
}

func (p Person) getFullName() string {
	return p.firstName + " " + p.lastName
}
