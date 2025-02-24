package main

import "fmt"

func main() {
	var b byte = 6
	var c byte
	c = 1
	b += 1
	d := 6
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	var x = []int{1, 2, 3}
	y := x
	y = append(x, 4)
	for i := range y {
		fmt.Println(y[i])
	}
}
