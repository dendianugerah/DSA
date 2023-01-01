package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")

	if len(words) != len(pattern) {
		return false
	}

	hashMap := make(map[interface{}]int)

	for i, v := range words {
		if hashMap[v] != hashMap[pattern[i]] {
			return false
		}
		hashMap[v] = i + 1
		hashMap[pattern[i]] = i + 1
	}

	return true
}

func main() {
	fmt.Println(wordPattern("abba", "cat dog dog cat"))
}
