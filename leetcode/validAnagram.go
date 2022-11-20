package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sM := make(map[rune]int)
	tM := make(map[rune]int)

	for _, v := range s {
		sM[v]++
	}
	for _, v := range t {
		tM[v]++
	}

	for k, v := range sM {
		if tM[k] != v {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
}
