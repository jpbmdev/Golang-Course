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

func (p *person) updateFirstName(firstName string) {
	p.firstName = firstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "alexmail@gmail.com",
			zipCode: 94000,
		},
	}

	alex.updateFirstName("carlos")

	alex.print()

}
