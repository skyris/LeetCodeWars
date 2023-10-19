package main

/*
Given an integer array nums, return the third distinct maximum number in this array. If the third maximum does not exist,
return the maximum number.
*/

import (
	"fmt"
	"math"
)

// Time Complexity: O(n)
// Space Complexity: O(1)
func thirdMax(nums []int) int {
	// it will work even if OS is 32-bit
	first := math.MinInt32
	second := math.MinInt32
	third := math.MinInt32
	diffNums := make(map[int]struct{})

	for i := 0; i < len(nums); i++ {
		if _, ok := diffNums[nums[i]]; !ok {
			diffNums[nums[i]] = struct{}{}
		}
		if nums[i] > first {
			first, second, third = nums[i], first, second
		} else if nums[i] > second && nums[i] < first {
			second, third = nums[i], second
		} else if nums[i] > third && nums[i] < second && nums[i] < first {
			third = nums[i]
		}
	}

	if len(diffNums) < 3 {
		return first
	}

	return third
}

func main() {
	nums := []int{1, 2, math.MinInt32}
	fmt.Println(thirdMax(nums))
}
