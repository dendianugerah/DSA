package main

func findGCD(nums []int) int {
	min := nums[0]
	max := nums[0]

	for _, val := range nums {
		if val > max {
			max = val
		}
		if val < min {
			min = val
		}
	}

	for i := min; i > 0; i-- {
		if max%i == 0 && min%i == 0 {
			return i
		}

	}
	return -1
}
