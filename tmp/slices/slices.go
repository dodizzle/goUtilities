package slices

import (
	"fmt"
	"slices"
)

func Slices() {
	var x []int
	x = append(x, 12, 23, 34)
	var y = []int{1, 2, 3}
	fmt.Println(slices.Equal(x, y))
	x = append(x, y...)
	x = append(x, 66)
	for i, v := range x {
		fmt.Printf("index=%d || value=%d \n", i, v)
	}
}
