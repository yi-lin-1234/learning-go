package main

import (
	"fmt"
)

type person struct {
	firstName   string
	lastName    string
	contactInfo // same as contactInfo contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Wang",
		contactInfo: contactInfo{
			email: "alex@gmail.com",
			zip:   11375,
		},
	}

	//alex.print()
	//alex.firstName = "John"
	//alexPointer := &alex           // memory address/pointer
	alex.updateName("John") // pass by value language so, it makes a copy of alex, therefore, it not work
	alex.print()

}

// receiver
func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName //* is like key, so you can access the value inside that address
}
