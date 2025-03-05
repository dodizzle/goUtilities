package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(FirstName string, LastName string, Age int) Person {
	p := Person{
		FirstName: FirstName,
		LastName:  LastName,
		Age:       Age,
	}
	return p
}

func MakePersonPointer(FirstName string, LastName string, Age int) *Person {
	p := Person{
		FirstName: FirstName,
		LastName:  LastName,
		Age:       Age,
	}
	return &p
}
func main() {
	p := MakePerson("John", "Doe", 25)
	p2 := MakePersonPointer("Jane", "Doe", 26)
	fmt.Println(p)
	fmt.Println(p2)

}
