package main

import "fmt"

func maximumWealth(account [][]int) int {
	result := 0
	for _, v := range account {
		max := 0
		for _, v2 := range v {
			max += v2
		}
		if result < max {
			result = max
		}
	}
	return result
}

func main() {
	fmt.Println(maximumWealth([][]int{{1, 2}, {2, 2}}))
}
