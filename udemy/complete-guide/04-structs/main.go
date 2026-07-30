package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateFirstName(name string) {
	p.firstName = name
}

func main() {
	bond := person{
		firstName: "James",
		lastName:  "Bond",
		contactInfo: contactInfo{
			email: "james@bond.co.uk",
			zip:   1000,
		},
	}
	bond.print()
	bond.updateFirstName("John")
	bond.contactInfo.email = "john@bond.co.uk"
	bond.print()

	name := "Bill"
	fmt.Println(&name)
}
