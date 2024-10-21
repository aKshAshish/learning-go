package main

import "fmt"

func Slice2D() {
	// Understanding 2D slices
	mat := [][]int{
		{1, 2, 3},
		{2},
		{4, 6},
	}
	fmt.Println(len(mat), cap(mat), mat) // 3 3 [[1 2 3] [2] [4 6]]
	// The append works similar to 1-D arrays
	mat = append(mat, []int{2, 9, 11})
	fmt.Println(len(mat), cap(mat), mat) // 4 6 [[1 2 3] [2] [4 6] [2 9 11]]

	fixedMat := [][3]int{
		{1, 2, 4},
		{2, 5, 6},
	}
	fmt.Println(len(fixedMat), cap(fixedMat), fixedMat) // 2 2 [[1 2 4] [2 5 6]]
	// The append works similar to 1-D arrays
	fixedMat = append(fixedMat, [3]int{2, 9, 11})
	fmt.Println(len(fixedMat), cap(fixedMat), fixedMat) // 3 4 [[1 2 4] [2 5 6] [2 9 11]]
}
