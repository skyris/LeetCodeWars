package main

/*
Given an array nums of integers, return how many of them contain an even number of digits.
*/

import (
	"fmt"
)

func findNumbers(nums []int) int {
	total := 0
	for _, n := range nums {
		if containsEvenNumOfDigits(n) {
			total += 1
		}
	}

	return total
}

func containsEvenNumOfDigits(n int) bool {
	if n == 0 {
		return false
	}

	count := 0
	for n > 0 {
		n = n / 10
		count += 1
	}

	if count%2 == 0 {
		return true
	}
	return false
}

func main() {
	nums := []int{555, 901, 482, 1771}
	fmt.Println(findNumbers(nums))
}
