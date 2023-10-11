package main

/*
Given an integer array nums sorted in non-decreasing order, return an array of the squares of each number sorted in non-decreasing order.
*/

import (
	"fmt"
)

// Two pointers method: O(n) and additional space
func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))
	left, right := 0, len(nums)-1
	index := len(nums) - 1

	for left <= right {
		leftSquare := nums[left] * nums[left]
		rightSquare := nums[right] * nums[right]

		if leftSquare > rightSquare {
			result[index] = leftSquare
			left++
		} else {
			result[index] = rightSquare
			right--
		}

		index--
	}

	return result
}

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(nums))
}
