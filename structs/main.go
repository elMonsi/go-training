package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	// using another struct
	contact contactInfo
}

func main() {
	// var alex person
	//
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)
	fmt.Printf("%+v", alex)

	alex.firstName = "Alejandro"
	alex.lastName = "Martinez"
	alex.contact.email = "test@test.com"
	alex.contact.zipCode = 75093

	fmt.Println(alex)
	fmt.Printf("%+v", alex)
}
