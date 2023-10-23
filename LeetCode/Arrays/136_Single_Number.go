package main

/*
Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
You must implement a solution with a linear runtime complexity and use only constant extra space.
*/

import (
	"fmt"
)

// Time complexity: O(n)
// Space complexity: O(1)
func singleNumber(nums []int) int {
	accum := 0
	for _, num := range nums {
		accum ^= num
	}

	return accum
}

func main() {
	nums := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(nums))
}
