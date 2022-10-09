package main

import "fmt"

func maxSubArray(nums []int) int {
	temp := 0
	res := nums[0]

	for _, value := range nums {
		if temp < value {
			temp += value
			if res < temp {
				res = temp
			}
		} else {
			if res < value {
				res = value
			}
		}
	}
	return temp
}

func main() {
	fmt.Println(maxSubArray([]int{1, 2, 3, 0, 4, 2}))
}
