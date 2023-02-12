package main

import "fmt"

func findLucky(arr []int) int {
	temp := make(map[int]int)

	for _, v := range arr {
		temp[v]++
	}

	oke := false
	largest := temp[0]
	for _, v := range temp {
		if v == temp[v] {
			current := temp[v]
			if current > largest {
				largest = current
				oke = true
			}
		}
	}
	if oke == true {
		return largest
	}
	return -1
}

func main() {
	fmt.Println(findLucky([]int{1, 2, 2, 3, 3, 3}))
}
