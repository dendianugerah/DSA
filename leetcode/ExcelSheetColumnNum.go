package main

import "fmt"

func titleToNumber(columnTitle string) int {
	var res int
	for _, v := range columnTitle {
		res *= 26
		res += (int(v) - 'A') + 1
	}
	return res
}

func main() {
	fmt.Println(titleToNumber("AB"))
}
