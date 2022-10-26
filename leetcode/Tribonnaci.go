package main

import "fmt"

func tribonnaci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}

	a, b, c := 0, 1, 1
	for i := 3; i <= n; i++ {
		temp := a + b + c
		a = b
		b = c
		c = temp
	}

	return c
}

func main() {
	fmt.Println(tribonnaci(25))
}
