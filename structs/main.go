package main

import (
	"fmt"
)

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
	josh := person{
		firstName: "Josh",
		lastName:  "Shih",
		contactInfo: contactInfo{
			email:   "josh@shih.com",
			zipCode: 12345,
		},
	}

	josh.updateName("Kablam")
	josh.print()
}

func (p person) print() {
	fmt.Println("First Name  :", p.firstName)
	fmt.Println("Last Name   :", p.lastName)
	fmt.Println("Contact Info:")
	fmt.Println("\tEmail  :", p.contactInfo.email)
	fmt.Println("\tZipCode:", p.contactInfo.zipCode)
}

func (p *person) updateName(newName string) {
	p.firstName = newName
}
