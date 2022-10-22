package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	// embedding structs
	// using a struct within another
	contact contactInfo
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (personPointer *person) updateName2(newFirstName string) {
	(*personPointer).firstName = newFirstName
}

func main() {
	// var alex person
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contact: contactInfo{
			email:   "mx@aol.com",
			zipCode: 12345,
		},
	}
	//fmt.Println(alex)
	fmt.Printf("%+v\n", alex)

	alex.firstName = "Alejandro"
	alex.lastName = "Martinez"
	alex.contact.email = "test@test.com"
	alex.contact.zipCode = 75093

	//fmt.Println(alex)
	alex.print()

	//This does not work
	alex.updateName("Hector")
	alex.print()

	//Do this instead
	personPointer := &alex
	personPointer.updateName2("Hector")
	alex.print()

	//Short way...
	//only works because the receiver (updateName2)
	//is told to use the memory address
	alex.updateName2("Victor")
	alex.print()
}
