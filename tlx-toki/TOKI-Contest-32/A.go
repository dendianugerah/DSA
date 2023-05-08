package main

import "fmt"

func main() {
	var n, a, b, c int

	fmt.Scan(&n)
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	isJumpYear := false

	if (n%a == 0 && n%b != 0) || n%c == 0 {
		isJumpYear = true
	}

	if isJumpYear {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
