package main

import (
	"fmt"
	"sort"
)

func trimMean(arr []int) float64 {
	sort.Ints(arr)
	sum, fivePercent := 0, len(arr)/20
	for i := fivePercent; i < len(arr)-fivePercent; i++ {
		sum += arr[i]
	}

	return float64(sum) / float64(len(arr)-fivePercent*2)
}

func main() {
	fmt.Println(trimMean([]int{1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3}))
}
