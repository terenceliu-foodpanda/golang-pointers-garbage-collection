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
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	//Version1: Original
	jimPointer := &jim
	jimPointer.updateName("jimmy1")
	jim.print()

	//Version2: Shortcut
	jim.updateName("jimmy2")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (pointerToPerson person) print() {
	fmt.Printf("\n%+v\n", pointerToPerson)
}
