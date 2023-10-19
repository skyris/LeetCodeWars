package main

/*
Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
Note that you must do this in-place without making a copy of the array.
*/

import (
	"fmt"
)

// Time Complexity: O(n)
// Space Complexity: O(1)
func moveZeroes(nums []int) {
	shift := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			shift++
			continue
		}
		nums[i-shift] = nums[i]
	}

	for i := len(nums) - 1; i >= len(nums)-shift; i-- {
		nums[i] = 0
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	fmt.Println(nums)
	moveZeroes(nums)
	fmt.Println(nums)
}
