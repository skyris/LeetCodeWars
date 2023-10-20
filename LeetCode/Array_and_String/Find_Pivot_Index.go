package main

/*
Given an array of integers nums, calculate the pivot index of this array.
The pivot index is the index where the sum of all the numbers strictly to the left of the index is equal
to the sum of all the numbers strictly to the index's right.
If the index is on the left edge of the array, then the left sum is 0 because there are no elements to the left.
This also applies to the right edge of the array.
Return the leftmost pivot index. If no such index exists, return -1.
1 <= nums.length <= 104
-1000 <= nums[i] <= 1000
*/

import (
	"fmt"
)

// Time complexity: O(n)
// Space complexity: O(1)
func pivotIndex(nums []int) int {
	if len(nums) < 1 {
		return -1
	}

	leftSum := 0
	rightSum := 0
	for i := 1; i < len(nums); i++ {
		rightSum += nums[i]
	}
	// First index
	if leftSum == rightSum {
		return 0
	}

	// Other indexes
	for i := 1; i < len(nums); i++ {
		leftSum += nums[i-1]
		rightSum -= nums[i]
		if leftSum == rightSum {
			return i
		}
	}

	return -1
}

func main() {
	// arr := []int{1, 7, 3, 6, 5, 6}
	// arr := []int{0, 1, -1}
	// arr := []int{0, 2}
	arr := []int{0}
	// arr := []int{-1, -1, 0, 1, 1, 0}
	fmt.Println(pivotIndex(arr))
}
