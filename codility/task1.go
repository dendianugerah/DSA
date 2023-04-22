package main

import (
	"fmt"
	"strconv"
)

func Helper(n int) string {
	return strconv.FormatInt(int64(n), 2)
}

func Solution(n int) int {
	temp := Helper(n)
	sl := []string{}
	result := 0
	current := 0
	highest := 0
	simp := 0

	for _, val := range temp {
		sl = append(sl, string(val))
	}

	for i := 0; i < len(sl); i++ {
		for j := i + 1; j < len(sl); j++ {
			if sl[j] == "1" {
				simp++
			}
		}

		if simp == 0 {
			return 0
		}

		if sl[i] != "1" {
			highest++
			current = highest
			if current >= result {
				result = current
			}
		} else {
			highest = 0
		}
	}

	return result

}

func main() {
	fmt.Println(Solution(1041))
}
