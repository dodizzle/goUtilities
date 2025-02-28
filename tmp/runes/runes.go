package runes

import "fmt"

func Runes() {
	type Employee struct {
		firstname string
		lastname  string
		id        int
	}
	person1 := Employee{
		id: 45,
	}
	person2 := Employee{
		firstname: "Tom",
		lastname:  "Brady",
	}
	var person3 Employee
	person3.firstname = "jo"
	person3.lastname = "blow"
	person3.id = 99
	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(person3)
}
