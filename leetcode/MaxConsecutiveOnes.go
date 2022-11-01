package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	var result int
	var temp int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			temp++
		}
		if temp > result {
			result = temp
		}
		if nums[i] == 0 {
			temp = 0
		}
	}
	return result
}

func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{1, 0, 1, 1, 0, 1}))
}
