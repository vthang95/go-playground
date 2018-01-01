package main

import "fmt"

type contactInfo struct {
	email       string
	phoneNumber string
}

type person struct {
	firstname string
	lastname  string
	contact   contactInfo
}

func main() {
	mommy := person{
		firstname: "Alex",
		lastname:  "Adison",
		contact: contactInfo{
			email:       "vthang95@gmail.com",
			phoneNumber: "0945143189",
		},
	}

	mommy.updateName("Huong")
	mommy.print()
}

func (pointerToPerson *person) updateName(newName string) {
	(*pointerToPerson).firstname = newName
}

func (pointerToPerson *person) print() {
	fmt.Printf("%+v", (*pointerToPerson))
}

// basic types + structs must use pointer to update value.
