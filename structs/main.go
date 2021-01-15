package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

//
func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func main() {

	jim := person{
		firstName: "Jim",
		lastName:  "Beasley",
		contact: contactInfo{
			email: "jim@jimbo.com",
			zip:   14214,
		},
	}

	// fmt.Printf("+%v", jim)

	jimPointer := &jim //has the address

	// fmt.Printf("the type %T \n", jimPointer)

	jimPointer.updateName("Jimmy") //address is transmitted
	jim.print()

}
