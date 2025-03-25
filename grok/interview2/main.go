package main

import "fmt"

// take in a slice of ints, sum them up one index value at a time
// and return the sum as an int
func sum(nums []int) int {
	var out int
	for _, v := range nums {
		if v%2 == 0 {
			out = out + v
		}
	}
	return out
}

func main() {
	// create slice of ints
	nums := []int{1, 2, 3, 4, 5, 6}
	// pass slice to sum function and return an int
	out := sum(nums)
	// Print out the sum of slice values
	fmt.Println(out)
}
