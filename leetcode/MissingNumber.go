package main

import (
	"fmt"
	"sort"
)

func missingNumber(nums []int) int {
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if i == len(nums) || i != nums[i] {
			return i
		}
	}

	return 0
}

func main() {
	fmt.Println(missingNumber([]int{3, 0, 1}))
}
