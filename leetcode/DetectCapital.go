package main

import (
	"fmt"
)

func isUpper(c byte) bool {
	return c >= 'A' && c <= 'Z'
}

func detectCapitalUse(word string) bool {
	count := 0
	for i := range word {
		if isUpper(word[i]) {
			count++
		}
	}

	return count == len(word) || count == 0 || (count == 1 && isUpper(word[0]))
}

func main() {
	fmt.Println(detectCapitalUse("USA"))
	fmt.Println(detectCapitalUse("FlaG"))
}
