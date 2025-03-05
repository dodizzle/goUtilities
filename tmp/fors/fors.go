package fors

import (
	"fmt"
	"math/rand"
)

func Fors() {
	// result := []int{}
	for range 100 {
		new := rand.Intn(100)
		// result = append(result, new)
		switch {
		case new%2 == 0:
			fmt.Printf("%d is divisible by 2\n", new)
		case new%3 == 0:
			fmt.Printf("%d is divisible by 3\n", new)
		case new%2 == 0 && new%3 == 0:
			fmt.Printf("%d is divisible by 2 & 3\n", new)
		default:
			fmt.Println("Nevermind!")
		}
	}
}
