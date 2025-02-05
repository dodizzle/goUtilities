package main

import (
	"math"
)

func closestToZero(nums []int) int {
	if nums == nil {
		return -1
	}
	if len(nums) == 0 {
		return 0
	}

	closest := nums[0]
	for _, num := range nums {
		if positiveValue(num) < positiveValue(closest) || (positiveValue(num) == positiveValue(closest) && num > closest) {
			closest = num
		}
	}
	return closest
}

func positiveValue(num int) int {
	return int(math.Abs(float64(num)))
}
