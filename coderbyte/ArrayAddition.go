package main

import "fmt"

func Sum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}

	return sum
}

func ArrayAddition(arr []int) bool {
	max := arr[0]

	for _, v := range arr {
		if v > max {
			max = v
		}
	}

	for i, v := range arr {
		if v == max {
			arr = append(arr[:i], arr[i+1:]...)
		}
	}

	var temp []int

	for i := 0; i < len(arr); i++ {
		temp = append(temp, arr[i])
		for j := i + 1; j < len(arr); j++ {
			temp = append(temp, arr[j])
			if Sum(temp) == max {
				return true
			}
		}
		temp = []int{}
	}

	return false
}

func main() {
	fmt.Println(ArrayAddition([]int{3, 5, -1, 8, 12}))
	fmt.Println(ArrayAddition([]int{5, 7, 16, 1, 2}))
}
