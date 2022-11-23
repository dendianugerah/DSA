package main

import "fmt"

func sortArrayByParity(nums []int) []int {
	numsIndex := 0

	for i := range nums {
		if nums[i]%2 == 0 {
			nums[i], nums[numsIndex] = nums[numsIndex], nums[i]
			numsIndex++
		}
	}

	// Solution 2 -- bruteforce
	// for i := 0; i < len(nums)-1; i++ {
	// 	for j := i + 1; j < len(nums); j++ {
	// 		if nums[j]%2 == 0 {
	// 			temp := nums[i]
	// 			nums[i] = nums[j]
	// 			nums[j] = temp
	// 		}
	// 	}
	// }

	return nums
}

func main() {
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4}))
}
