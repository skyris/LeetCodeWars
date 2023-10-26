package main

/*
Given an m x n matrix, return all elements of the matrix in spiral order.
*/

import (
	"fmt"
)

// Time complexity O(n*m)
// Space complexity O(n*m)
func spiralOrder(matrix [][]int) []int {
	w, h := len(matrix[0]), len(matrix)

	xMin, xMax := 0, w-1
	yMin, yMax := 0, h-1
	index := 0
	x, y := 0, 0
	res := make([]int, w*h)

	for {
		// Right
		if index >= w*h {
			break
		}
		for x = xMin; x <= xMax; x++ {
			res[index] = matrix[y][x]
			index++
		}
		x--
		yMin++

		// Down
		if index >= w*h {
			break
		}
		for y = yMin; y <= yMax; y++ {
			res[index] = matrix[y][x]
			index++
		}
		y--
		xMax--

		// Left
		if index >= w*h {
			break
		}
		for x = xMax; x >= xMin; x-- {
			res[index] = matrix[y][x]
			index++
		}
		x++
		yMax--

		// Up
		if index >= w*h {
			break
		}
		for y = yMax; y >= yMin; y-- {
			res[index] = matrix[y][x]
			index++
		}
		y++
		xMin++
	}

	return res
}

func main() {
	mat := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(spiralOrder(mat))
}
