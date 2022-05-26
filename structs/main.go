package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func main() {
	alex := person{"Alex", "Anderson", contactInfo{"alex.anderson@mail.me", 12000}}

	james := person{firstName: "James", lastName: "Johnson"}

	var lily person
	lily.firstName = "Lily"
	lily.lastName = "Collins"

	jim := person{
		firstName: "Jim",
		lastName:  "Carry",
		contact: contactInfo{
			email:   "jim.carrry@mail.me",
			zipCode: 46900,
		},
	}

	// jimPointer := &jim
	// jimPointer.updateName("Jimmy")
	jim.updateName("Jimmy")

	alex.print()
	james.print()
	lily.print()
	jim.print()
}
