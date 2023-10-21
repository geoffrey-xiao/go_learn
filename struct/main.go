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

func main() {
	// geoffrey := person{
	// 	firstName: "Geoffrey",
	// 	lastName:  "Xiao",
	// }
	// var geoffrey person
	// geoffrey.firstName = "Geoffrey"
	// geoffrey.lastName = "Xiao"
	geoffrey := person{
		firstName: "Geoffrey",
		lastName:  "Xiao",
		contact: contactInfo{
			email:   "geoffrey@gmail.com",
			zipCode: 0000,
		},
	}
	// fmt.Println(geoffrey)
	// fmt.Printf("%+v", geoffrey)
	// geoffreyPointer := &geoffrey
	// geoffreyPointer.updateName("jeff")
	geoffrey.firstName = "Summer"
	geoffrey.print()
}

func (pointerToPerson *person) updateName(newName string) {
	(*pointerToPerson).firstName = newName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
