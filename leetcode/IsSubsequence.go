package main

import (
	"fmt"
)

func isSubsequence(s, t string) bool {
	var index int

	for x := range t {
		if index == len(s) {
			return true
		}
		if t[x] == s[index] {
			index++
		}
	}

	return index == len(s)
}

func main() {

	fmt.Println(isSubsequence("abc", "ahbgdc"))
}
