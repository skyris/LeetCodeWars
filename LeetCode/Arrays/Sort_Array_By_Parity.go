package main

/*
Given an integer array nums, move all the even integers at the beginning of the array followed by all the odd integers.
Return any array that satisfies this condition.
*/

import (
	"fmt"
)

// Time Complexity: O(n)
// Space Complexity: O(1)
func sortArrayByParity(nums []int) []int {
	shift := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 != 0 {
			shift++
			continue
		}
		nums[i-shift], nums[i] = nums[i], nums[i-shift]
	}

	return nums
}

func main() {
	nums := []int{3, 1, 2, 4}
	fmt.Println(sortArrayByParity(nums))
}
