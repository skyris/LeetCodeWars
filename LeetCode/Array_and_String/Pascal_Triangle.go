package main

/*
Given an integer numRows, return the first numRows of Pascal's triangle.
Example 1:
Input: numRows = 5
Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
*/

import (
	"fmt"
)

func generate(numRows int) [][]int {
	res := make([][]int, numRows)

	res[0] = []int{1}
	for y := 1; y < numRows; y++ {
		layer := make([]int, y+1)
		layer[0], layer[y] = 1, 1
		for x := 1; x < y; x++ {
			layer[x] = res[y-1][x-1] + res[y-1][x]
		}
		res[y] = layer
	}

	return res
}

func main() {
	fmt.Println(generate(5))
}
