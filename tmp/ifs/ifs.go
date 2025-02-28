package ifs

import "fmt"

func Ifs() {
	var x []int
	x = append(x, 12, 23, 34)
	n := 0
	count := len(x)
	for n <= count {
		fmt.Println(count, x[n])
		n++
	}

}
