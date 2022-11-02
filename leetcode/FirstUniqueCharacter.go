package main

import "fmt"

func firstUniqChar(s string) byte {
	m := make(map[rune]int)

	for _, val := range s {
		m[val]++
	}

	for _, c := range s {
		if m[c] == 2 {
			return byte(c)
		}
	}
	return 0
}

func main() {
	fmt.Println(firstUniqChar("abccbaacz"))
}
