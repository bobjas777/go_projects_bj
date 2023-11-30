package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode string
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jimmy",
		lastName:  "Logan",
		contactInfo: contactInfo{
			email:   "jim@example.com",
			zipCode: "C4F6H8",
		},
	}

	jim.updateName("John")
	jim.print()
}
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
func (p person) print() {
	fmt.Printf("%+v", p)
}
