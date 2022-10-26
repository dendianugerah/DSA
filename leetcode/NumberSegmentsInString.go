package main

import "fmt"

func countSegments(s string) int {
	if s == "" {
		return 0
	}

	res := 1

	for i := 0; i < len(s); i++ {
		if s[i] == 32 && s[i-1] != 32 {
			res++
		}
	}
	return res
}

func main() {
	fmt.Println(countSegments("Hello, my name is John"))
}
