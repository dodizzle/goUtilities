package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct {
	sound string
}

type Cat struct {
	sound string
}

func (d Dog) Speak() string {
	return d.sound
}

func (c Cat) Speak() string {
	return c.sound
}

func makeSound(s Speaker) {
	fmt.Println(s.Speak())
}
func main() {
	c := Cat{"Meow"}
	d := Dog{"Woof"}
	makeSound(c)
	makeSound(d)
}
