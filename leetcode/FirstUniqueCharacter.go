package main

import "fmt"

func firstUniqChar(s string) int {
	m := make(map[rune]int)

	for _, val := range s {
		m[val]++
	}

	for i, c := range s {
		if m[c] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(firstUniqChar("loveleetcode"))
}
