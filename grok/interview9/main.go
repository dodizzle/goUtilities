package main

import (
	"fmt"
	"os"
)

func checkEmpty(usage []int) bool {
	if len(usage) == 0 {
		return false
	}
	return true
}

func getAverage(usage []int, length int) float64 {
	sum := 0
	for _, v := range usage {
		sum = sum + v
	}
	return (float64(sum) / float64(length))
}

func main() {
	//Write a Go function that takes a slice of integers representing memory usage (in MB)
	//  and returns the average usage as a float64.
	// Handle the case where the slice is empty by returning 0.0.
	usage := []int{512, 256, 768, 1024}
	if !checkEmpty(usage) {
		fmt.Println("No data!")
		os.Exit(0)
	}
	average := getAverage(usage, len(usage))

	fmt.Println(average)
}
