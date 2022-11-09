package main

import "fmt"

func numIdenticalPairs(nums []int) int {
	var result int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] && i < j {
				result++
			}
		}
	}
	return result
}

func main() {
	fmt.Println(numIdenticalPairs([]int{1, 2, 3, 1, 1, 3}))
}
