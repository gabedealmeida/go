package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	bob := person{
		firstName: "Bob",
		lastName:  "The Builder",
		contact: contactInfo{
			email:   "bob@thebuilder.com",
			zipcode: 00000,
		},
	}

	bobPointer := &bob
	bobPointer.updateName("Bobby")
	bob.print() // it now prints the update struct with the firstName Bobby

}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
