package main

/*
Given a fixed-length integer array arr, duplicate each occurrence of zero, shifting the remaining elements to the right.
Note that elements beyond the length of the original array are not written.
Do the above modifications to the input array in place and do not return anything.
*/

import (
	"fmt"
)

// O(n)
func duplicateZeros(arr []int) {
	zeros, length := 0, len(arr)-1

	for i := 0; i <= length-zeros; i++ {
		if arr[i] == 0 {
			if i == length-zeros {
				arr[length] = 0
				length--
				break
			}
			zeros++
		}
	}

	for i := length - zeros; i >= 0; i-- {
		if arr[i] == 0 {
			arr[i+zeros] = 0
			zeros--
			arr[i+zeros] = 0
		} else {
			arr[i+zeros] = arr[i]
		}
	}
}

// O(n^2)
func duplicateZerosNaive(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == 0 {
			copy(arr[i+1:], arr[i:])
			arr[i+1] = 0
			i++
		}
	}
}

func main() {
	arr := []int{1, 0, 2, 3, 0, 4, 5, 0}
	fmt.Println(arr)
	duplicateZeros(arr)
	fmt.Println(arr)
}
