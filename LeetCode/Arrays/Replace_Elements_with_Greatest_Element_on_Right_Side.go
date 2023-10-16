package main

/*
Given an array arr, replace every element in that array with the greatest element among the elements to its right,
and replace the last element with -1.
After doing so, return the array.
*/

import (
	"fmt"
)

// Time complexity: O(n)
// Space complexity: O(1)
func replaceElements(arr []int) []int {
	nextMax := -1

	for i := len(arr) - 1; i > 0; i-- {
		arr[i], nextMax = nextMax, max(arr[i], nextMax)
	}

	return arr
}

func main() {
	arr := []int{17, 18, 5, 4, 6, 1}
	fmt.Println(replaceElements(arr))
}
