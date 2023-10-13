package main

/*
Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once.
The relative order of the elements should be kept the same. Then return the number of unique elements in nums.
*/

import (
	"fmt"
)

// Time complexity O(n)
// Space complexity O(1)
func removeDuplicates(nums []int) int {
	shift := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			shift++
		}
		nums[i-shift] = nums[i]
	}

	return len(nums) - shift
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(nums)
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}
