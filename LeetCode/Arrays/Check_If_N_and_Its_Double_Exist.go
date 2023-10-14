package main

/*
Given an array arr of integers, check if there exist two indices i and j such that :

i != j
0 <= i, j < arr.length
arr[i] == 2 * arr[j]
*/

import (
	"fmt"
)

// Time complexity: O(n)
// Space complexity: O(n)
func checkIfExist(arr []int) bool {
	seen := make(map[int]struct{})

	for _, v := range arr {
		_, ok := seen[v]
		if ok {
			return true
		} else {
			seen[2*v] = struct{}{}
			if v%2 == 0 {
				seen[v/2] = struct{}{}
			}
		}
	}

	return false
}

func main() {
	arr := []int{4, -7, 11, 4, 18}
	fmt.Println(checkIfExist(arr))
}
