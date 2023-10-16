package main

/*
Given an array of integers arr, return true if and only if it is a valid mountain array.

Recall that arr is a mountain array if and only if:

arr.length >= 3
There exists some i with 0 < i < arr.length - 1 such that:
arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
arr[i] > arr[i + 1] > ... > arr[arr.length - 1]
*/

import (
	"fmt"
)

// Time complexity: O(n)
// Space complexity: O(1)
func validMountainArray(arr []int) bool {
	length := len(arr)
	if length < 3 {
		return false
	}

	peak := 0
	for i := 1; i < length; i++ {
		if arr[i] == arr[i-1] {
			return false
		}
		if arr[i] < arr[i-1] {
			break
		}
		peak = i
	}

	if peak == 0 || peak == length-1 {
		return false
	}

	for i := peak + 1; i < length; i++ {
		if arr[i] >= arr[i-1] {
			return false
		}
	}

	return true
}

func main() {
	arr := []int{3, 5, 5}
	fmt.Println(validMountainArray(arr))
}
