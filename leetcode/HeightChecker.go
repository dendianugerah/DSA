package main

import "fmt"

func heightChecker(heights []int) int {
	result := 0
	l := []int{}

	for _, v := range heights {
		l = append(l, v)
	}

	for i := 0; i < len(heights); i++ {
		for j := i + 1; j < len(heights); j++ {
			if heights[i] > heights[j] {
				temp := heights[i]
				heights[i] = heights[j]
				heights[j] = temp
			}
		}
	}

	for i, v := range heights {
		if v != l[i] {
			result++
		}
	}

	return result
}

func main() {
	fmt.Println(heightChecker([]int{1, 1, 4, 2, 1, 3}))
}
