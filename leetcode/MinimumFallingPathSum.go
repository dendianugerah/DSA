package main

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)

	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			minPrev := matrix[i-1][j]

			if j > 0 && matrix[i-1][j-1] < minPrev {
				minPrev = matrix[i-1][j-1]
			}

			if j < n-1 && matrix[i-1][j+1] < minPrev {
				minPrev = matrix[i-1][j+1]
			}

			matrix[i][j] += minPrev
		}
	}

	out := matrix[n-1][0]
	for i := 1; i < n; i++ {
		if matrix[n-1][i] < out {
			out = matrix[n-1][i]
		}
	}

	return out
}
