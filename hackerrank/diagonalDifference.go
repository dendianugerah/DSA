package main

import (
	"math"
)

func diagonalDifference(arr [][]int32) int32 {
	// Write your code here
	var result int32
	var leftDiagonal int32
	var rightDiagonal int32

	for i := 0; i < len(arr); i++ {
		leftDiagonal += arr[i][i]
		rightDiagonal += arr[i][len(arr)-1-i]
	}

	result = leftDiagonal - rightDiagonal

	return int32(math.Abs(float64(result)))
}
