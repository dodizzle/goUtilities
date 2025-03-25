package main

import (
	"github.com/dodizzle/goUtilities/interview1/work"
)

// create a function that takes a list of integers and returns the integer closest to 0
// if the closest int is negative and positive, return the positive one
// if the list is empty, return 0

func main() {
	list := []int{6, 10, -2, 2, 3, 1, 4, 5}

	// find the int closest to 0
	result := work.ClosestToZero(list)
	println(result)
}
