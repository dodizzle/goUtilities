package main

// create a function that takes a list of integers and returns the integer closest to 0
// if the closest int is negative and positive, return the positive one
// if the list is empty, return 0

func main() {
	list := []int{10, -2, 2, 3, 4, 5}

	// find the int closest to 0
	result := closestToZero(list)
	println(result)
}
