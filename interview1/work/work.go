package work

import (
	"math"
)

func ClosestToZero(nums []int) int {
	if nums == nil {
		return -1
	}
	if len(nums) == 0 {
		return 0
	}

	closest := nums[0]
	for _, num := range nums {
		if PositiveValue(num) < PositiveValue(closest) || (PositiveValue(num) == PositiveValue(closest) && num > closest) {
			closest = num
		}
	}
	return closest
}

func PositiveValue(num int) int {
	return int(math.Abs(float64(num)))
}
