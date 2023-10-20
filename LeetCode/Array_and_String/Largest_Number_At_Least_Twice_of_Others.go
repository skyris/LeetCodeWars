package main

/*
You are given an integer array nums where the largest integer is unique.
Determine whether the largest element in the array is at least twice as much as every other number in the array.
If it is, return the index of the largest element, or return -1 otherwise.
2 <= nums.length <= 50
0 <= nums[i] <= 100
The largest element in nums is unique.
*/

import (
	"fmt"
)

// Time complexity: O(n)
// Space complexity: O(1)
func dominantIndex(nums []int) int {
	first, second := -1, -1
	index := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] > first {
			first, second = nums[i], first
			index = i
		}
		if nums[i] > second && nums[i] < first {
			second = nums[i]
		}
	}

	if first/2 >= second {
		return index
	}

	return -1
}

func main() {
	nums := []int{3, 6, 1, 0}
	fmt.Println(dominantIndex(nums))
}
