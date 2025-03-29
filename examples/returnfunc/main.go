package main

import "fmt"

func prefixer(first string) func(string) string {
	return func(ret string) string {
		return first + " " + ret
	}

}

func main() {

	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))
}
