package main

/*
Given an m x n matrix mat, return an array of all the elements of the array in a diagonal order.
*/

import (
	"fmt"
)

// Time complexity O(n*m)
// Space complexity O(n*m)
func findDiagonalOrder(mat [][]int) []int {
	w, h := len(mat[0]), len(mat)
	res := []int{}

	for s := 0; s < w+h; s++ {
		x, y := initCoords(s, w, h)
		for x >= 0 && x < w && y >= 0 && y < h {
			res = append(res, mat[y][x])
			if s%2 == 0 { // Odd - Up
				x++
				y--
			} else { // Even - down
				x--
				y++
			}
		}
	}

	return res
}

func initCoords(s int, w int, h int) (int, int) {
	var x, y int
	/* Sum of both coords equals s */
	if s%2 == 0 {
		// Odd - Up
		if s < h {
			y = s
		} else {
			y = h - 1
		}
		x = s - y

		return x, y
	}
	// Even - down
	if s < w {
		x = s
	} else {
		x = w - 1
	}
	y = s - x

	return x, y
}

func main() {
	mat := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10, 11, 12},
	}

	fmt.Println(findDiagonalOrder(mat))
}
