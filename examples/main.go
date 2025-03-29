package main

import "fmt"

type Humans struct {
	Name   string
	Age    int
	Gender string
	Measurements
}

type Measurements struct {
	Height int
	Weight int
}

// standard function
// func editHuman(human *Humans) {
// 	human.Name = "Doe"
// 	human.Age = 30
// 	human.Gender = "Female"
// 	human.Height = 99
// }

// method
// func (human Humans) editHuman(newName string) string {
// 	human.Name = newName
// 	return fmt.Sprintf("name:%s || age:%d || gender:%s || height:%d", human.Name, human.Age, human.Gender, human.Height)
// }

type Foo struct {
	alpha int
	beta  int
}

type Answer struct {
	answer int
}

func (f Foo) Add() (a Answer) {
	a.answer = f.alpha + f.beta
	return a
}

func main() {
	foo := Foo{alpha: 2, beta: 6}
	returnedAnswer := foo.Add()
	fmt.Println(returnedAnswer)
	// human := Humans{
	// 	Name:   "John",
	// 	Age:    25,
	// 	Gender: "Male",
	// }
	// fmt.Println(human)
	// f := human.editHuman("Sally")
	// fmt.Println(f)
	// mm := []string{"a", "b", "c"}
	// for _, v := range mm {
	// 	fmt.Println(v)
	// }

}
