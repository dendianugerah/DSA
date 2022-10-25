package main

import "fmt"

func divisorGame(n int) bool {
	if n%2 == 0 {
		return true
	}

	return false
}

func main() {
	fmt.Println(divisorGame(3))
}
