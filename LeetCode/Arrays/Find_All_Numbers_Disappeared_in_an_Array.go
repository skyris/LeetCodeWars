package main

/*
Given an array nums of n integers where nums[i] is in the range [1, n], return an array of all the integers
in the range [1, n] that do not appear in nums.
*/

import (
	"fmt"
)

// Time Complexity: O(n)
// Space Complexity: O(1)
func findDisappearedNumbers(nums []int) (result []int) {
	for _, num := range nums {
		if nums[abs(num)-1] > 0 {
			nums[abs(num)-1] *= -1
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			result = append(result, i+1)
		}
	}

	return result
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers(nums))
}
