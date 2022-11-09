package main

import "fmt"

func areOccurrencesEqual(s string) bool {
	m := make(map[rune]int)

	for _, v := range s {
		m[v]++
	}

	res := []int{}
	for _, v := range m {
		res = append(res, v)
	}

	for i := 0; i < len(res)-1; i++ {
		if res[i] != res[i+1] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(areOccurrencesEqual("abacbc"))
}
