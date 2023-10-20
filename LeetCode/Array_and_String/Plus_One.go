package main

/*
You are given a large integer represented as an integer array digits, where each digits[i] is the ith digit of the integer.
The digits are ordered from most significant to least significant in left-to-right order. The large integer does not contain any leading 0's.
Increment the large integer by one and return the resulting array of digits.
1 <= digits.length <= 100
0 <= digits[i] <= 9
digits does not contain any leading 0's.
*/

import (
	"fmt"
)

// Time complexity: O(n)
// Space complexity: O(n)
func plusOne(digits []int) []int {
	res := make([]int, len(digits))
	prev := 1
	for i := len(digits) - 1; i >= 0; i-- {
		res[i] = (digits[i] + prev) % 10
		prev = (digits[i] + prev) / 10
	}

	if prev == 1 {
		res = append([]int{1}, res...)
	}

	return res
}

// Time complexity: O(n)
// Space complexity: O(1)
func plusOne2(digits []int) []int {
Loop:
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += 1
		switch {
		case digits[i] <= 9:
			break Loop
		case i > 0:
			digits[i] = 0
		default:
			digits[i] = 0
			digits = append([]int{1}, digits...)
		}
	}

	return digits
}

func main() {
	digit := []int{9, 9, 9}
	fmt.Println(plusOne(digit))
}
