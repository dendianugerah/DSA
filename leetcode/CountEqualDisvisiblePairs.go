package main

import "fmt"

func countPairs(nums []int, k int) int {
	var result int

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] && (i*j)%k == 0 {
				result++
			}
		}
	}

	return result
}

func main() {
	fmt.Print(countPairs([]int{3, 1, 2, 2, 2, 1, 3}, 2))
}
