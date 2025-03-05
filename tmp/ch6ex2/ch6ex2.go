package main

import "fmt"

func UpdateSlice(pizza []string, piece string) {
	pizza[len(pizza)-1] = piece
	fmt.Printf("Before UpdateSlice:%s\n", pizza)
}

func GrowSlice(pizza []string, piece string) {
	pizza = append(pizza, piece)
	fmt.Printf("Before GlowSlice:%s\n", pizza)
}

func main() {
	pizza := make([]string, 3)
	pizza = append(pizza, "a", "b", "c")
	UpdateSlice(pizza, "d")
	fmt.Printf("After UpdateSlice:%s\n", pizza)
	GrowSlice(pizza, "e")
	fmt.Printf("After GlowSlice:%s\n", pizza)

}
