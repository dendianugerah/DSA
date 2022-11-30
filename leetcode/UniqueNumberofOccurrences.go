package main

import "fmt"

func uniqueOccurrences(arr []int) bool {
	ArrMap := make(map[int]int)

	for _, v := range arr {
		ArrMap[v]++
	}

	storedArr := []int{}
	for _, v := range ArrMap {
		storedArr = append(storedArr, v)
	}

	for i := 0; i < len(storedArr)-1; i++ {
		for j := i + 1; j < len(storedArr); j++ {
			if storedArr[i] == storedArr[j] {
				return false
			}
		}
	}

	return true
}

func main() {
	arr := []int{1, 2}
	fmt.Println(uniqueOccurrences(arr))
}
