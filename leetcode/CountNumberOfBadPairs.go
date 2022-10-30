package main

import "fmt"

func countBadPairs(nums []int) int64 {
	var res int64

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if j-i != nums[j]-nums[i] && i < j {
				res++
			}
		}
	}

	return res
}

func main() {
	fmt.Print(countBadPairs([]int{4, 1, 3, 3}))
}
