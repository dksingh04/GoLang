package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {

	Deepak := person{
		firstName: "Deepak",
		lastName:  "Singh",
		contactInfo: contactInfo{
			email: "a@b.com",
			zip:   95051,
		},
	}
	//fmt.Println(Deepak)

	Deepak.print()

	// other way of intantiating struct

	var deepak person
	deepak.firstName = "Deepak"
	deepak.lastName = "Singh"
	fmt.Println(deepak)
	fmt.Printf("%+v", deepak)

}

func (p person) print() {
	fmt.Printf("%+v", p)
}
