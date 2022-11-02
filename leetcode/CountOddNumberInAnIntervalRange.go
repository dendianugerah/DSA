package main

import "fmt"

func countOdds(low int, high int) int {
	res := 0
	for i := low; i <= high; i++ {
		if i%2 != 0 {
			res++
		}
	}
	return res
}

func main() {
	fmt.Println(countOdds(3, 7))
}
