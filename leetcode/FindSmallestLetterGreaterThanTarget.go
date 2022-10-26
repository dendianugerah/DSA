package main

import "fmt"

func nextGreatestLetter(letters []byte, target byte) byte {
	for i := 0; i < len(letters); i++ {
		if target < letters[i] {
			return letters[i]
		}
	}
	return letters[0]
}

func main() {
	fmt.Println(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'a'))
}
