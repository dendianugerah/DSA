package main

import "fmt"

func search(nums []int, target int) int {
	down := 0
	up := len(nums) - 1

	for down <= up {
		mid := (down + up) / 2

		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			up = mid - 1
		} else {
			down = mid + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(search([]int{1, 2, 3, 4, 5, 6}, 2))
}
