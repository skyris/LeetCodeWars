package main

/*
Given a binary array nums, return the maximum number of consecutive 1's in the array.

Example 1:

Input: nums = [1,1,0,1,1,1]
Output: 3
Explanation: The first two digits or the last three digits are consecutive 1s. The maximum number of consecutive 1s is 3.
Example 2:

Input: nums = [1,0,1,1,0,1]
Output: 2

Constraints:

1 <= nums.length <= 105
nums[i] is either 0 or 1.
*/

import (
	"fmt"
)

func findMaxConsecutiveOnes(nums []int) int {
	result := 0
	current := 0

	for _, val := range nums {
		if val == 0 {
			current = 0
			continue
		}

		current += 1
		result = max(result, current)
	}

	return result
}

func main() {
	nums := []int{1, 1, 0, 1, 1, 1}
	fmt.Println(findMaxConsecutiveOnes(nums))
}
