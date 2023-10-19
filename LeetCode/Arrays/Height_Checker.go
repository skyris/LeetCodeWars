package main

/*
A school is trying to take an annual photo of all the students. The students are asked to stand in a single file line
in non-decreasing order by height. Let this ordering be represented by the integer array expected where expected[i]
is the expected height of the ith student in line.
You are given an integer array heights representing the current order that the students are standing in.
Each heights[i] is the height of the ith student in line (0-indexed).
Return the number of indices where heights[i] != expected[i].
*/

import (
	"fmt"
	"sort"
)

// Time Complexity: O(n)
// Space Complexity: O(n)
func heightChecker(heights []int) int {
	expected := make([]int, len(heights))
	copy(expected, heights)
	sort.Ints(expected)
	count := 0
	for i := 0; i < len(heights); i++ {
		if heights[i] != expected[i] {
			count++
		}
	}

	return count
}

func main() {
	heights := []int{1, 1, 4, 2, 1, 3}
	fmt.Println(heightChecker(heights))
}
