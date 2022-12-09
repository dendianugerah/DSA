package main

import (
	"fmt"
)

func distributeCandies(candyType []int) int {

	res, max := 0, len(candyType)/2

	m := make(map[int]bool)
	for _, v := range candyType {
		if _, ok := m[v]; !ok {
			m[v] = true
			res++
			if res == max {
				return max
			}
		}
	}

	return res

}

func main() {
	fmt.Println(distributeCandies([]int{1, 1, 2, 2, 3, 3}))
}
