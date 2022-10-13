package main

import "fmt"

func moveZeroes(nums []int) {
	ans := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[ans], nums[i] = nums[i], nums[ans]
			ans++
		}
	}
	fmt.Println(nums)
}

func main() {
	moveZeroes([]int{1, 2, 0, 2, 0, 2})
}
