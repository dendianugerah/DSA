package main

import (
	"fmt"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	Minimu := arr[1] - arr[0]
	result := [][]int{}
	for i := 1; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] < Minimu {
			Minimu = arr[i+1] - arr[i]
		}
	}

	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] == Minimu {
			result = append(result, []int{arr[i-1], arr[i]})
		}
	}
	return result
}

func main() {
	fmt.Println(minimumAbsDifference([]int{4, 2, 1, 3}))
}
