package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	alex := person{
		firstName: "Jacek",
		lastName:  "Dupa",
		contactInfo: contactInfo{
			email:   "dupa@gmail.com",
			zipCode: 43430,
		},
	}

	alex.updateName("Tomek")
	alex.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
