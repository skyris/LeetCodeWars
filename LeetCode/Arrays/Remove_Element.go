package main

/*
Given an integer array nums and an integer val, remove all occurrences of val in nums in-place.
The order of the elements may be changed. Then return the number of elements in nums which are not equal to val.
*/

import (
	"fmt"
)

// Time complexity O(n)
// Space complexity O(1)
func removeElement(nums []int, val int) int {
	shift := 0
	length := len(nums)
	for i := 0; i < length; i++ {
		if nums[i] == val {
			shift++
			continue
		}
		nums[i-shift] = nums[i]
	}

	return length - shift
}

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	fmt.Println(removeElement(nums, val))
	fmt.Println(nums)
}
